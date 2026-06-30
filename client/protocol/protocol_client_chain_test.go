package protocol

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/flare-foundation/flare-system-client/utils"
)

// recorder collects the order in which chain steps run.
type recorder struct {
	mu   sync.Mutex
	runs []string
}

func (r *recorder) record(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.runs = append(r.runs, name)
}

func (r *recorder) snapshot() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]string, len(r.runs))
	copy(out, r.runs)
	return out
}

func TestRunChain_RunsAllStepsInOrder(t *testing.T) {
	rec := &recorder{}
	steps := []submitterStep{
		{offset: 0, run: func(context.Context) { rec.record("a") }},
		{offset: 5 * time.Millisecond, run: func(context.Context) { rec.record("b") }},
		{offset: 10 * time.Millisecond, run: func(context.Context) { rec.record("c") }},
	}
	runChain(context.Background(), steps)

	got := rec.snapshot()
	want := []string{"a", "b", "c"}
	if len(got) != len(want) {
		t.Fatalf("ran %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("ran %v, want %v", got, want)
		}
	}
}

func TestRunChain_CancelledBeforeFirstStepRunsNothing(t *testing.T) {
	rec := &recorder{}
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // already cancelled

	steps := []submitterStep{
		{offset: 50 * time.Millisecond, run: func(context.Context) { rec.record("a") }},
		{offset: 60 * time.Millisecond, run: func(context.Context) { rec.record("b") }},
	}
	runChain(ctx, steps)

	if got := rec.snapshot(); len(got) != 0 {
		t.Fatalf("expected no steps to run, ran %v", got)
	}
}

// The core invariant: once the first step has run, cancellation must not skip
// the remaining steps (a submit1 commit obligates its submit2 reveal).
func TestRunChain_ObligationContinuesAfterCancelDuringFirstStep(t *testing.T) {
	rec := &recorder{}
	ctx, cancel := context.WithCancel(context.Background())

	steps := []submitterStep{
		{offset: 0, run: func(context.Context) {
			rec.record("first")
			cancel() // shutdown requested right after the first step ran
		}},
		{offset: 5 * time.Millisecond, run: func(context.Context) { rec.record("second") }},
		{offset: 10 * time.Millisecond, run: func(context.Context) { rec.record("third") }},
	}
	runChain(ctx, steps)

	got := rec.snapshot()
	want := []string{"first", "second", "third"}
	if len(got) != len(want) {
		t.Fatalf("obligation not honored: ran %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("obligation not honored: ran %v, want %v", got, want)
		}
	}
}

// The obligation steps' work (RunEpoch -> GetPayload -> submit) honors its
// context and no-ops once cancelled, so post-gate steps must receive a LIVE
// (detached) context, otherwise the submission silently does nothing and the
// round is penalised anyway.
func TestRunChain_ObligationStepsRunUnderLiveContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	var mu sync.Mutex
	var ranCancelled []string
	checkLive := func(name string, c context.Context) {
		if c.Err() != nil {
			mu.Lock()
			ranCancelled = append(ranCancelled, name)
			mu.Unlock()
		}
	}

	steps := []submitterStep{
		{offset: 0, run: func(c context.Context) {
			checkLive("first", c)
			cancel() // shutdown requested right after the first step (submit1) ran
		}},
		{offset: 5 * time.Millisecond, run: func(c context.Context) { checkLive("second", c) }},
		{offset: 10 * time.Millisecond, run: func(c context.Context) { checkLive("third", c) }},
	}
	runChain(ctx, steps)

	mu.Lock()
	defer mu.Unlock()
	if len(ranCancelled) > 0 {
		t.Fatalf("obligation steps ran under a cancelled context %v; the real submitters "+
			"would skip submission and the round would still be penalised", ranCancelled)
	}
}

// Each step's context must be cancelled once the step returns, so anything it
// spawned is released rather than dangling.
func TestRunChain_StepContextCancelledAfterStepReturns(t *testing.T) {
	var mu sync.Mutex
	var captured []context.Context
	record := func(c context.Context) {
		mu.Lock()
		defer mu.Unlock()
		captured = append(captured, c)
	}
	steps := []submitterStep{
		{offset: 0, run: func(c context.Context) { record(c) }},
		{offset: 5 * time.Millisecond, run: func(c context.Context) { record(c) }},
	}
	runChain(context.Background(), steps)

	mu.Lock()
	defer mu.Unlock()
	if len(captured) != 2 {
		t.Fatalf("expected 2 steps to run, got %d", len(captured))
	}
	for i, c := range captured {
		if c.Err() == nil {
			t.Errorf("step %d context should be cancelled after the step returns", i)
		}
	}
}

// A panic in one step must not prevent the remaining steps from running: a
// submit2 that blows up still leaves submitSignatures owed (and vice versa).
func TestRunChain_PanicInStepDoesNotStopChain(t *testing.T) {
	rec := &recorder{}
	steps := []submitterStep{
		{offset: 0, run: func(context.Context) { rec.record("first") }},
		{offset: 2 * time.Millisecond, run: func(context.Context) { panic("boom") }},
		{offset: 4 * time.Millisecond, run: func(context.Context) { rec.record("third") }},
	}
	runChain(context.Background(), steps)

	got := rec.snapshot()
	want := []string{"first", "third"}
	if len(got) != len(want) {
		t.Fatalf("panic stopped the chain: ran %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("panic stopped the chain: ran %v, want %v", got, want)
		}
	}
}

func TestSubmitterChain_Composition(t *testing.T) {
	period := 90 * time.Second
	ticker := &utils.EpochTicker{Epoch: &utils.EpochTimingConfig{Period: period}}
	const epoch int64 = 7

	submitter := func(off time.Duration) *Submitter {
		return &Submitter{SubmitterBase: SubmitterBase{startOffset: off}}
	}
	sigSubmitter := func(off time.Duration) *SignatureSubmitter {
		return &SignatureSubmitter{SubmitterBase: SubmitterBase{startOffset: off}}
	}

	t.Run("all_enabled", func(t *testing.T) {
		c := &client{
			submitter1:         submitter(1 * time.Second),
			submitter2:         submitter(2 * time.Second),
			signatureSubmitter: sigSubmitter(3 * time.Second),
		}
		chain := c.submitterChain(ticker, epoch)
		wantOffsets := []time.Duration{1 * time.Second, period + 2*time.Second, period + 3*time.Second}
		if len(chain) != len(wantOffsets) {
			t.Fatalf("want chain of %d, got %d", len(wantOffsets), len(chain))
		}
		for i, s := range chain {
			if s.offset != wantOffsets[i] {
				t.Errorf("step %d offset = %v, want %v", i, s.offset, wantOffsets[i])
			}
		}
	})

	t.Run("fdc_tail_only", func(t *testing.T) {
		c := &client{
			submitter2:         submitter(2 * time.Second),
			signatureSubmitter: sigSubmitter(3 * time.Second),
		}
		if chain := c.submitterChain(ticker, epoch); len(chain) != 2 {
			t.Fatalf("want chain of 2, got %d", len(chain))
		}
	})

	t.Run("none_enabled", func(t *testing.T) {
		c := &client{}
		if chain := c.submitterChain(ticker, epoch); len(chain) != 0 {
			t.Fatalf("want empty chain, got %d", len(chain))
		}
	})
}

// FDC edge: a two-step chain (submit2 gate -> submitSignatures). Once submit2
// has run, submitSignatures must run and be waited for even if shutdown is
// requested in between, or FDC is penalised.
func TestRunChain_Submit2ObligesSignatures(t *testing.T) {
	rec := &recorder{}
	ctx, cancel := context.WithCancel(context.Background())

	steps := []submitterStep{
		{offset: 0, run: func(context.Context) {
			rec.record("submit2")
			cancel() // shutdown right after submit2 ran
		}},
		{offset: 5 * time.Millisecond, run: func(context.Context) { rec.record("submitSignatures") }},
	}
	runChain(ctx, steps)

	got := rec.snapshot()
	if len(got) != 2 || got[0] != "submit2" || got[1] != "submitSignatures" {
		t.Fatalf("submit2 must oblige submitSignatures despite cancellation: ran %v", got)
	}
}

// Post-gate obligation steps run concurrently: a slow step must not delay the
// others (under a sequential chain "fast" would record after "slow").
func TestRunChain_ObligationsRunConcurrently(t *testing.T) {
	rec := &recorder{}
	steps := []submitterStep{
		{offset: 0, run: func(context.Context) { rec.record("gate") }},
		{offset: 0, run: func(context.Context) { time.Sleep(150 * time.Millisecond); rec.record("slow") }},
		{offset: 0, run: func(context.Context) { rec.record("fast") }},
	}
	runChain(context.Background(), steps)

	got := rec.snapshot()
	idx := func(name string) int {
		for i, n := range got {
			if n == name {
				return i
			}
		}
		return -1
	}
	if len(got) != 3 {
		t.Fatalf("expected all 3 steps to run, got %v", got)
	}
	if idx("fast") > idx("slow") {
		t.Errorf("obligations did not run concurrently: %v (fast should precede slow)", got)
	}
}

// A slow gate must not delay the start of later steps: under a sequential chain
// the later step would only start after the gate returned; here it starts on its
// own offset, concurrently.
func TestRunChain_SlowGateDoesNotDelayLaterSteps(t *testing.T) {
	rec := &recorder{}
	steps := []submitterStep{
		{offset: 0, run: func(context.Context) {
			time.Sleep(150 * time.Millisecond)
			rec.record("gate")
		}},
		{offset: 10 * time.Millisecond, run: func(context.Context) { rec.record("later") }},
	}
	runChain(context.Background(), steps)

	got := rec.snapshot()
	if len(got) != 2 || got[0] != "later" || got[1] != "gate" {
		t.Fatalf("later step blocked on slow gate: ran %v (want [later gate])", got)
	}
}

package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	clientContext "github.com/flare-foundation/flare-system-client/client/context"
	"github.com/flare-foundation/flare-system-client/client/runner"
	"github.com/flare-foundation/flare-system-client/client/shared"
	globalConfig "github.com/flare-foundation/flare-system-client/config"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

func main() {
	logger.Info("Starting Flare System client")

	clientCtx, err := clientContext.BuildContext()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	logger.Set(clientCtx.Config().Logger)

	for _, w := range clientCtx.Config().SubmitterWarnings() {
		logger.Warnf("submitter config: %s", w)
	}

	for _, w := range clientCtx.Config().GasOverrideWarnings() {
		logger.Warnf("gas config: %s", w)
	}

	// txs are signed with the configured chain_id — a mismatch fails every send
	verifyCtx, cancelVerify := context.WithTimeout(context.Background(), 10*time.Second)
	err = clientCtx.Config().Chain.VerifyChainID(verifyCtx)
	cancelVerify()
	if errors.Is(err, globalConfig.ErrChainIDMismatch) {
		logger.Fatalf("chain config: %v", err)
	} else if err != nil {
		logger.Warnf("chain config: could not verify chain_id against the node: %v", err)
	}

	// a half-filled cutover entry would silently keep the old Relay past the switch
	if err := shared.ValidateRelayCutovers(); err != nil {
		logger.Fatalf("relay config: %v", err)
	}

	// Prometheus metrics
	shared.InitMetricsServer(&clientCtx.Config().Metrics)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sig := <-signalChan
		logger.Infof("Received %v signal, attempting graceful shutdown", sig)
		cancel()
	}()

	wg := runner.Start(ctx, cancel, clientCtx)
	wg.Wait()
	logger.Info("Stopped Flare System client")
}

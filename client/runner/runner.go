package runner

import (
	"context"
	"errors"
	"reflect"
	"sync"

	clientContext "github.com/flare-foundation/flare-system-client/client/context"
	"github.com/flare-foundation/flare-system-client/client/epoch"
	"github.com/flare-foundation/flare-system-client/client/finalizer"
	"github.com/flare-foundation/flare-system-client/client/protocol"
	"github.com/flare-foundation/flare-system-client/client/shared"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type Runner interface {
	Run(ctx context.Context) error
}

// RunAsync runs a runner in a go routine with context and adds it to work group.
func RunAsync(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup, r Runner) {
	if r == nil || reflect.ValueOf(r).IsNil() {
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := r.Run(ctx)
		if err != nil && !errors.Is(err, context.Canceled) {
			logger.Errorf("Unexpected error, terminating: %v", err)
			cancel()
		}
	}()
}

// Start sets up registrationClient, protocolClient, and finalizerClient, then asynchronously runs all of them and returns their workgroup.
func Start(ctx context.Context, cancel context.CancelFunc, clientCtx clientContext.ClientContext) *sync.WaitGroup {
	registrationClient, err := epoch.NewClient(clientCtx)
	if err != nil {
		logger.Fatalf("Error creating registration client: %v", err)
	}

	messageChannel := make(chan shared.ProtocolMessage, 2*len(clientCtx.Config().Protocol)) // twice just to be on the save side

	protocolClient, err := protocol.NewClient(clientCtx, messageChannel)
	if err != nil {
		logger.Fatalf("Error creating protocol client: %v", err)
	}
	finalizerClient, err := finalizer.NewClient(clientCtx, messageChannel)
	if err != nil {
		logger.Fatalf("Error creating finalizer client: %v", err)
	}

	wg := sync.WaitGroup{}
	RunAsync(ctx, cancel, &wg, protocolClient)
	RunAsync(ctx, cancel, &wg, registrationClient)
	RunAsync(ctx, cancel, &wg, finalizerClient)

	return &wg
}

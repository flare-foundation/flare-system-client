package runner

import (
	"context"
	"errors"
	clientContext "flare-fsc/client/context"
	"flare-fsc/client/epoch"
	"flare-fsc/client/finalizer"
	"flare-fsc/client/protocol"
	"flare-fsc/client/shared"
	"flare-fsc/logger"
	"reflect"
	"sync"
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
			logger.Error("Unexpected error, terminating: %v", err)
			cancel()
		}
	}()
}

// Start sets up registrationClient, protocolClient, and finalizerClient, then asynchronously runs all of them and returns their workgroup.
func Start(ctx context.Context, cancel context.CancelFunc, clientCtx clientContext.ClientContext) *sync.WaitGroup {
	registrationClient, err := epoch.NewEpochClient(clientCtx)
	if err != nil {
		logger.Fatal("Error creating registration client: %v", err)
	}

	messageChannel := make(chan shared.ProtocolMessage, 2*len(clientCtx.Config().Protocol)) // twice just to be on the save side

	protocolClient, err := protocol.NewProtocolClient(clientCtx, messageChannel)
	if err != nil {
		logger.Fatal("Error creating protocol client: %v", err)
	}
	finalizerClient, err := finalizer.NewFinalizerClient(clientCtx, messageChannel)
	if err != nil {
		logger.Fatal("Error creating finalizer client: %v", err)
	}

	wg := sync.WaitGroup{}
	RunAsync(ctx, cancel, &wg, protocolClient)
	RunAsync(ctx, cancel, &wg, registrationClient)
	RunAsync(ctx, cancel, &wg, finalizerClient)

	return &wg
}

package runner

import (
	"context"
	"errors"
	clientContext "flare-tlc/client/context"
	"flare-tlc/client/finalizer"
	"flare-tlc/client/protocol"
	"flare-tlc/client/registration"
	"flare-tlc/logger"
	"reflect"
	"sync"
)

type Runner interface {
	Run(ctx context.Context) error
}

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

func Start(ctx context.Context, cancel context.CancelFunc, clientCtx clientContext.ClientContext) *sync.WaitGroup {
	registrationClient, err := registration.NewRegistrationClient(clientCtx)
	if err != nil {
		logger.Fatal("Error creating registration client: %v", err)
	}
	protocolClient, err := protocol.NewProtocolClient(clientCtx)
	if err != nil {
		logger.Fatal("Error creating protocol client: %v", err)
	}
	finalizerClient, err := finalizer.NewFinalizerClient(clientCtx)
	if err != nil {
		logger.Fatal("Error creating finalizer client: %v", err)
	}

	wg := sync.WaitGroup{}
	RunAsync(ctx, cancel, &wg, protocolClient)
	RunAsync(ctx, cancel, &wg, registrationClient)
	RunAsync(ctx, cancel, &wg, finalizerClient)

	return &wg
}

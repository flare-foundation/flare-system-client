package main

import (
	"context"
	clientContext "flare-fsc/client/context"
	"flare-fsc/client/runner"
	"flare-fsc/client/shared"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

func main() {
	logger.Info("Starting Flare System client")

	clientCtx, err := clientContext.BuildContext()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
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

package main

import (
	"context"
	clientContext "flare-fsc/client/context"
	"flare-fsc/client/runner"
	"flare-fsc/client/shared"
	"flare-fsc/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
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
		logger.Info("Received %v signal, attempting graceful shutdown", sig)
		cancel()
	}()

	wg := runner.Start(ctx, cancel, clientCtx)
	wg.Wait()
	logger.Info("Stopped Flare System client")
}

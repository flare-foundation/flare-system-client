package main

import (
	"context"
	clientContext "flare-tlc/client/context"
	"flare-tlc/client/runner"
	"flare-tlc/client/shared"
	"flare-tlc/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger.Info("Starting flare top level client")

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
		logger.Info("Received signal, stopping: %v", sig)
		cancel()
	}()

	wg := runner.Start(ctx, cancel, clientCtx)
	wg.Wait()
	logger.Info("Stopped flare top level client")
}

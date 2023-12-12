package main

import (
	"flare-tlc/client/context"
	"flare-tlc/client/migrations"
	"flare-tlc/client/runner"
	"flare-tlc/client/shared"
	"flare-tlc/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, err := context.BuildContext()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	err = migrations.Container.ExecuteAll(ctx.DB())
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, os.Interrupt, syscall.SIGTERM)

	// Prometheus metrics
	shared.InitMetricsServer(&ctx.Config().Metrics)

	runner.Start(ctx)

	<-cancelChan
	logger.Info("Stopped flare top level client")

}

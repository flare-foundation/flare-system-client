package runner

import (
	"flare-tlc/client/clients"
	"flare-tlc/client/context"
)

func Start(ctx context.ClientContext) {
	client, err := clients.NewVotingClient(ctx.Config())
	if err != nil {
		panic(err)
	}
	go client.Run()
}

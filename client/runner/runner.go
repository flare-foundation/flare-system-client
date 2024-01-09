package runner

import (
	"flare-tlc/client/clients"
	"flare-tlc/client/context"
)

func Start(ctx context.ClientContext) {
	client, err := clients.NewVotingClient(ctx)
	if err != nil {
		panic(err)
	}
	go func() {
		err := client.Run()
		if err != nil {
			panic(err)
		}
	}()
}

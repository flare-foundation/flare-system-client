package runner

import (
	"flare-tlc/client/clients"
	"flare-tlc/client/context"
)

func Start(ctx context.ClientContext) {
	client := clients.NewVotingClient()
	go client.Run()
}

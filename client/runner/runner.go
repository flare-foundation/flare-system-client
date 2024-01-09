package runner

import (
	"flare-tlc/client/clients"
	"flare-tlc/client/context"
)

func Start(ctx context.ClientContext) {
	// votingClient, err := clients.NewVotingClient(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	registrationClient, err := clients.NewRegistratinClient(ctx)
	if err != nil {
		panic(err)
	}

	// go votingClient.Run()
	go func() {
		err := registrationClient.Run()
		if err != nil {
			panic(err)
		}
	}()

}

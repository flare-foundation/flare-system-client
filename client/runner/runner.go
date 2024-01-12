package runner

import (
	"flare-tlc/client/context"
	clients "flare-tlc/client/protocol"
	"flare-tlc/client/registration"
	"reflect"
)

type Runner interface {
	Run() error
}

func RunAsync(r Runner) {
	if r == nil || reflect.ValueOf(r).IsNil() {
		return
	}

	go func() {
		err := r.Run()
		if err != nil {
			panic(err)
		}
	}()
}

func Start(ctx context.ClientContext) {
	votingClient, err := clients.NewVotingClient(ctx)
	if err != nil {
		panic(err)
	}
	registrationClient, err := registration.NewRegistrationClient(ctx)
	if err != nil {
		panic(err)
	}
	RunAsync(votingClient)
	RunAsync(registrationClient)
}

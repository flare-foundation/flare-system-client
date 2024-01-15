package runner

import (
	"flare-tlc/client/context"
	"flare-tlc/client/protocol"
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
	protocolClient, err := protocol.NewProtocolClient(ctx)
	if err != nil {
		panic(err)
	}
	registrationClient, err := registration.NewRegistrationClient(ctx)
	if err != nil {
		panic(err)
	}
	RunAsync(protocolClient)
	RunAsync(registrationClient)
}

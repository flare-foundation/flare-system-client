package clients

import "flare-tlc/utils/contracts/registry"

type RegistryVoter struct {
	voterRegistry *registry.Registry
}

func NewRegistryVoter() *RegistryVoter {
	return &RegistryVoter{}
}

func (v *RegistryVoter) RegisterVoter() <-chan bool {

	return nil
}

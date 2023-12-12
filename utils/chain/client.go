package chain

import (
	"github.com/ava-labs/avalanchego/utils/rpc"
)

func ClientOptions(apiKey string) []rpc.Option {
	if len(apiKey) == 0 {
		return []rpc.Option{}
	} else {
		return []rpc.Option{rpc.WithQueryParam("x-apikey", apiKey)}
	}
}

func RPCClientOptions(apiKey string) string {
	if len(apiKey) == 0 {
		return ""
	} else {
		return "?x-apikey=" + apiKey
	}
}

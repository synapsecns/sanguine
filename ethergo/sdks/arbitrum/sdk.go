package arbitrum

import "github.com/ethereum/go-ethereum/accounts/abi/bind"

type ArbitrumSDK interface {
	// TODO:
}

type arbitrumSDKImpl struct {
	client bind.ContractBackend
}

func NewArbitrumSDK(client bind.ContractBackend) {

}

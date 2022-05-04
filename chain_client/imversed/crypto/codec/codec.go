package codec

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	ethsecp256k12 "github.com/imversed/relayer/chain_client/imversed/crypto/ethsecp256k1"
)

// RegisterInterfaces register the Ethermint key concrete types.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k12.PubKey{})
}

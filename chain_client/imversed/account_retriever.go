package client

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/imversed/imversed/x/evm/types"
)

// EnsureExists returns an error if no account exists for the given address else nil.
func (cc *ChainClient) EnsureExists(clientCtx client.Context, addr sdk.AccAddress) error {
	if _, _, err := cc.GetAccountNumberSequence(clientCtx, addr); err != nil {
		return err
	}

	return nil
}

// GetAccountNumberSequence returns sequence and account number for the given address.
// It returns an error if the account couldn't be retrieved from the state.
func (cc *ChainClient) GetAccountNumberSequence(clientCtx client.Context, addr sdk.AccAddress) (uint64, uint64, error) {
	var header metadata.MD
	address := ethereum.BytesToAddress(addr).String()

	queryClient := types.NewQueryClient(cc)
	res, err := queryClient.CosmosAccount(context.Background(), &types.QueryCosmosAccountRequest{Address: address}, grpc.Header(&header))
	if err != nil {
		return 0, 0, err
	}

	return res.AccountNumber, res.Sequence, nil
}

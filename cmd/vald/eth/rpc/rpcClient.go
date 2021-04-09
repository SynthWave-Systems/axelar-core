package rpc

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//go:generate moq -out ./mock/rpcClient.go -pkg mock . Client

// Client provides calls to an Ethereum RPC endpoint
type Client interface {
	BlockNumber(ctx context.Context) (uint64, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// ClientImpl implements Client
type ClientImpl struct {
	*ethclient.Client
}

// NewClient returns an Ethereum rpc client
func NewClient(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	// try to access ethereum network
	if _, err := client.ChainID(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}

// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package client

import (
	"context"
	"strings"

	"github.com/AnomalyFi/seq-sdk/requester"
	"github.com/AnomalyFi/seq-sdk/types"
)

type JSONRPCClient struct {
	requester *requester.EndpointRequester

	networkID uint32
	chainID   string
}

// New creates a new client object.
func NewJSONRPCClient(uri string, networkID uint32, chainID string) *JSONRPCClient {
	uri = strings.TrimSuffix(uri, "/")
	uri += "/tokenapi"
	req := requester.New(uri, "tokenvm")
	return &JSONRPCClient{
		requester: req,
		networkID: networkID,
		chainID:   chainID,
	}
}

type SubmitMsgTxArgs struct {
	ChainId          string `json:"chain_id"`
	NetworkID        uint32 `json:"network_id"`
	SecondaryChainId []byte `json:"secondary_chain_id"`
	Data             []byte `json:"data"`
}

type SubmitMsgTxReply struct {
	TxID string `json:"txId"`
}

func (cli *JSONRPCClient) SubmitTx(
	ctx context.Context,
	ChainId string,
	NetworkID uint32,
	SecondaryChainId []byte,
	Data []byte,
) (string, error) {
	resp := new(SubmitMsgTxReply)
	err := cli.requester.SendRequest(
		ctx,
		"submitMsgTx",
		&SubmitMsgTxArgs{
			ChainId:          ChainId,
			NetworkID:        NetworkID,
			SecondaryChainId: SecondaryChainId,
			Data:             Data,
		},
		resp,
	)
	return resp.TxID, err
}

func (cli *JSONRPCClient) GetBlockHeadersByHeight(
	ctx context.Context,
	height uint64,
	end int64,
) (*types.BlockHeadersResponse, error) {
	resp := new(types.BlockHeadersResponse)
	err := cli.requester.SendRequest(
		ctx,
		"getBlockHeadersByHeight",
		&types.GetBlockHeadersByHeightArgs{
			Height: height,
			End:    end,
		},
		resp,
	)
	return resp, err
}

func (cli *JSONRPCClient) GetBlockHeadersID(
	ctx context.Context,
	id string,
	end int64,
) (*types.BlockHeadersResponse, error) {
	resp := new(types.BlockHeadersResponse)
	err := cli.requester.SendRequest(
		ctx,
		"getBlockHeadersId",
		&types.GetBlockHeadersIDArgs{
			ID:  id,
			End: end,
		},
		resp,
	)
	return resp, err
}

func (cli *JSONRPCClient) GetBlockHeadersByStart(
	ctx context.Context,
	start int64,
	end int64,
) (*types.BlockHeadersResponse, error) {
	resp := new(types.BlockHeadersResponse)
	err := cli.requester.SendRequest(
		ctx,
		"getBlockHeadersByStart",
		&types.GetBlockHeadersByStartArgs{
			Start: start,
			End:   end,
		},
		resp,
	)
	return resp, err
}

// func (cli *JSONRPCClient) GetBlockTransactions(
// 	ctx context.Context,
// 	id string,
// ) (*types.TransactionResponse, error) {
// 	resp := new(types.TransactionResponse)
// 	//TODO does this need to be lowercase for the string?
// 	err := cli.requester.SendRequest(
// 		ctx,
// 		"getBlockTransactions",
// 		&types.GetBlockTransactionsArgs{
// 			ID: id,
// 		},
// 		resp,
// 	)
// 	return resp, err
// }

func (cli *JSONRPCClient) GetBlockTransactionsByNamespace(
	ctx context.Context,
	height uint64,
	namespace string,
) (*types.SEQTransactionResponse, error) {
	resp := new(types.SEQTransactionResponse)
	err := cli.requester.SendRequest(
		ctx,
		"getBlockTransactionsByNamespace",
		&types.GetBlockTransactionsByNamespaceArgs{
			Height:    height,
			Namespace: namespace,
		},
		resp,
	)
	return resp, err
}

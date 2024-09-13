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
	uri += "/seqapi"
	req := requester.New(uri, "seqvm")
	return &JSONRPCClient{
		requester: req,
		networkID: networkID,
		chainID:   chainID,
	}
}

type SubmitMsgTxArgs struct {
	ChainID          string   `json:"chain_id"`
	NetworkID        uint32   `json:"network_id"`
	SecondaryChainID []byte   `json:"secondary_chain_id"`
	Data             [][]byte `json:"data"`
}

type SubmitMsgTxReply struct {
	TxID string `json:"txId"`
}

func (cli *JSONRPCClient) SubmitMsgTx(
	ctx context.Context,
	ChainId string,
	NetworkID uint32,
	SecondaryChainId []byte,
	Data [][]byte,
) (string, error) {
	resp := new(SubmitMsgTxReply)
	err := cli.requester.SendRequest(
		ctx,
		"submitMsgTx",
		&SubmitMsgTxArgs{
			ChainID:          ChainId,
			NetworkID:        NetworkID,
			SecondaryChainID: SecondaryChainId,
			Data:             Data,
		},
		resp,
	)
	return resp.TxID, err
}

func (cli *JSONRPCClient) GetBlockHeadersByHeight(
	ctx context.Context,
	height uint64,
	endTimeStamp int64,
) (*types.BlockHeadersResponse, error) {
	resp := new(types.BlockHeadersResponse)
	err := cli.requester.SendRequest(
		ctx,
		"getBlockHeadersByHeight",
		&types.GetBlockHeadersByHeightArgs{
			Height:       height,
			EndTimeStamp: endTimeStamp,
		},
		resp,
	)
	return resp, err
}

func (cli *JSONRPCClient) GetBlockHeadersID(
	ctx context.Context,
	id string,
	endTimeStamp int64,
) (*types.BlockHeadersResponse, error) {
	resp := new(types.BlockHeadersResponse)
	err := cli.requester.SendRequest(
		ctx,
		"getBlockHeadersId",
		&types.GetBlockHeadersIDArgs{
			ID:           id,
			EndTimeStamp: endTimeStamp,
		},
		resp,
	)
	return resp, err
}

func (cli *JSONRPCClient) GetBlockHeadersByStartTimeStamp(
	ctx context.Context,
	startTimeStamp int64,
	endTimeStamp int64,
) (*types.BlockHeadersResponse, error) {
	resp := new(types.BlockHeadersResponse)
	err := cli.requester.SendRequest(
		ctx,
		"getBlockHeadersByStartTimeStamp",
		&types.GetBlockHeadersByStartTimeStampArgs{
			StartTimeStamp: startTimeStamp,
			EndTimeStamp:   endTimeStamp,
		},
		resp,
	)
	return resp, err
}

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

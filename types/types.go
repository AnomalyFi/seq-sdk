package types

type BlockInfo struct {
	BlockId   string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	L1Head    uint64 `json:"l1_head"`
	Height    uint64 `json:"height"`
}

type BlockHeadersResponse struct {
	From   uint64      `json:"from"`
	Blocks []BlockInfo `json:"blocks"`
	Prev   BlockInfo   `json:"prev"`
	Next   BlockInfo   `json:"next"`
}

type SEQTransaction struct {
	Namespace   string `json:"namespace"`
	TxID        string `json:"tx_id"`
	Index       uint64 `json:"tx_index"`
	Transaction []byte `json:"transaction"`
}

type SEQTransactionResponse struct {
	Txs     []*SEQTransaction `json:"txs"`
	BlockID string            `json:"id"`
}

type GetBlockHeadersByHeightArgs struct {
	Height       uint64 `json:"height"`
	EndTimeStamp int64  `json:"end_timestamp"`
}

type GetBlockHeadersIDArgs struct {
	ID           string `json:"id"`
	EndTimeStamp int64  `json:"end_timestamp"`
}

type GetBlockHeadersByStartTimeStampArgs struct {
	StartTimeStamp int64 `json:"start_timestamp"`
	EndTimeStamp   int64 `json:"end_timestamp"`
}

type GetBlockTransactionsArgs struct {
	ID string `json:"block_id"`
}

type GetBlockTransactionsByNamespaceArgs struct {
	Height    uint64 `json:"height"`
	Namespace string `json:"namespace"`
}

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

// type TransactionResponse struct {
// 	Txs     []*chain.Transaction `json:"txs"`
// 	BlockId string               `json:"id"`
// }

type SEQTransaction struct {
	Namespace   string `json:"namespace"`
	Tx_id       string `json:"tx_id"`
	Index       uint64 `json:"tx_index"`
	Transaction []byte `json:"transaction"`
}

// TODO need to fix this. Tech debt
type SEQTransactionResponse struct {
	Txs       []*SEQTransaction `json:"txs"`
	BlockId   string            `json:"id"`
	Timestamp int64             `json:"timestamp"`
	L1Head    uint64            `json:"l1_head"`
	Height    uint64            `json:"height"`
}

type LastAcceptedReply struct {
	Height    uint64 `json:"height"`
	BlockID   string `json:"blockId"`
	Timestamp int64  `json:"timestamp"`
}

type GetBlockHeadersByHeightArgs struct {
	Height uint64 `json:"height"`
	End    int64  `json:"end"`
}

type GetBlockHeadersIDArgs struct {
	ID  string `json:"id"`
	End int64  `json:"end"`
}

type GetBlockHeadersByStartArgs struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type GetBlockTransactionsArgs struct {
	ID string `json:"block_id"`
}

type GetBlockTransactionsByNamespaceArgs struct {
	Height    uint64 `json:"height"`
	Namespace string `json:"namespace"`
}

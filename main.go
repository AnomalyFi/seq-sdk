package main

import (
	"context"
	"fmt"
	"time"

	trpc "github.com/AnomalyFi/seq-sdk/client"
)

// TODO really need
func FetchHeadersForWindow(ctx context.Context) error {

	id := "ah7Ngy4Cg7TJ6CnEzB3E8twiF7PmaLSeUYpLxxtrXhKtMBrcn"

	urlNew := "http://127.0.0.1:9654/ext/bc/ah7Ngy4Cg7TJ6CnEzB3E8twiF7PmaLSeUYpLxxtrXhKtMBrcn"

	cli := trpc.NewJSONRPCClient(urlNew, 1337, id)

	//TODO this is submitting a unix timestamp so I may have a mismatch because HyperSDK uses int64 for timestamp and this uses uint64 for timestamps
	start_time := time.Now().UnixMilli()
	end_time := time.Now().UnixMilli() - 3000000
	// start_time := start * 1000
	// end_time := end * 1000
	// start_time = 1725708078363
	res, err := cli.GetBlockHeadersByStartTimeStamp(context.Background(), start_time, end_time)
	if err != nil {
		return err
	}

	fmt.Println(res.Blocks)

	return nil
}

func main() {
	err := FetchHeadersForWindow(context.Background())
	if err != nil {
		fmt.Println("Error occured:", err)
	}
}

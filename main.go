package main

import (
	"context"
	"fmt"
	"time"

	trpc "github.com/AnomalyFi/seq-sdk/client"
)

// TODO really need
func FetchHeadersForWindow(ctx context.Context) error {

	id := "2qMoscnJNq7h9XkLzWBGdFmvSMhnctXfHbiifQSfNN7shyA8SR"

	urlNew := "http://127.0.0.1:41887/ext/bc/2qMoscnJNq7h9XkLzWBGdFmvSMhnctXfHbiifQSfNN7shyA8SR"

	cli := trpc.NewJSONRPCClient(urlNew, 1337, id)
	//var res WindowStart
	//getBlockHeadersByStart

	//TODO this is submitting a unix timestamp so I may have a mismatch because HyperSDK uses int64 for timestamp and this uses uint64 for timestamps
	start := time.Now().Unix()

	end := time.Now().Unix() - 120

	start_time := start * 1000

	end_time := end * 1000

	//start=1,698,195,498 end=1,698,195,500

	res, err := cli.GetBlockHeadersByStart(context.Background(), int64(start_time), int64(end_time))
	if err != nil {
		return err
	}

	fmt.Println(res.Blocks[0])

	return nil
}

func main() {
	fmt.Println("Hello World")
	err := FetchHeadersForWindow(context.Background())
	if err != nil {
		fmt.Println("Error occured:", err)
	}
}

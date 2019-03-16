package main

import (
	"context"
	"fmt"

	horizonclient "github.com/stellar/go/exp/clients/horizon"
)

func main() {
	c := horizonclient.DefaultTestNetClient

	er := horizonclient.EffectRequest{Limit: 10, Cursor: "now"}

	ctx := context.Background()

	fmt.Println("starting")

	err := c.Stream(ctx, er, func(data interface{}) {
		fmt.Println(data)
	})

	if err != nil {
		fmt.Printf("error: %s", err)
	}
}

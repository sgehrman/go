package main

import (
	"context"
	"encoding/json"
	"fmt"

	horizonclient "github.com/stellar/go/exp/clients/horizon"
)

func main() {
	c := horizonclient.DefaultTestNetClient

	er := horizonclient.EffectRequest{Cursor: "now"}

	ctx := context.Background()

	fmt.Println("EffectRequest stream test")

	err := c.Stream(ctx, er, func(data interface{}) {
		json, err := json.MarshalIndent(data, "", "    ")

		if err != nil {
			fmt.Println("error in MarshalIndent: ", err)
		}
		fmt.Println(string(json))
	})

	if err != nil {
		fmt.Printf("error: %s", err)
	}
}

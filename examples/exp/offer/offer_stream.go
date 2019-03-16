package main

import (
	"context"
	"encoding/json"
	"fmt"

	horizonclient "github.com/stellar/go/exp/clients/horizon"
)

func main() {
	c := horizonclient.DefaultTestNetClient

	request := horizonclient.OfferRequest{ForAccount: "GBEHNI5AKMOIMFXZE6YIWGO26ZNUZFYOGVZQYUAVTG72VREMF4FPSZ5A"}

	ctx := context.Background()

	fmt.Println("OfferRequest stream test")

	err := c.Stream(ctx, request, func(data interface{}) {
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

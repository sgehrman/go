package main

import (
	"context"
	"encoding/json"
	"fmt"

	horizonclient "github.com/stellar/go/exp/clients/horizon"
)

func main() {
	c := horizonclient.DefaultTestNetClient

	request := horizonclient.AssetRequest{ForAssetIssuer: "GCYF2F6QATA4L43DSXTL72D6BEBFKTN6OW7Y65PFS3QSPLDO5LOHSAYZ"}

	ctx := context.Background()

	fmt.Println("AssetRequest stream test")

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

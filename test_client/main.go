package main

import (
	"fmt"
	"github.com/firehydrant/api-client-go/test_client/client"
)

func main() {
	payload, err := client.Start()
	if err != nil {
		fmt.Println(
			"error: ",
			err,
		)
	}
	fmt.Println(
		"payload: ",
		payload,
	)
}

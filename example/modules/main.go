package main

import (
	"context"
	"fmt"

	"github.com/chelnak/go-forge-yourself/forge"
)

func main() {
	client := forge.NewClient()

	ctx := context.Background()

	opts := &forge.ListModulesOptions{
		Limit:        100,
		Endorsements: []forge.Endorsement{forge.EndorsementSupported},
		Owner:        "puppetlabs",
	}

	response, err := client.Modules.ListModules(ctx, opts)
	if err != nil {
		panic(err)
	}

	for _, module := range response.Results {
		fmt.Println(module.Slug)
	}
}

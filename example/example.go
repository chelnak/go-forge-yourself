package main

import (
	"context"
	"fmt"

	"github.com/chelnak/go-forge-yourself/forge"
)

func main() {
	client := forge.NewClient()

	ctx := context.Background()

	// ListModules
	listModuleOpts := &forge.ListModulesOptions{
		Limit:        100,
		Endorsements: []forge.Endorsement{forge.EndorsementSupported},
		Owner:        "puppetlabs",
	}

	listModulesResponse, err := client.Modules.ListModules(ctx, listModuleOpts)
	if err != nil {
		panic(err)
	}

	for _, module := range listModulesResponse.Results {
		fmt.Println(module.Slug)
	}

	fmt.Println("Number of modules with the supported endorsement:", len(listModulesResponse.Results))
}

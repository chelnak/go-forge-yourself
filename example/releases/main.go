package main

import (
	"context"
	"fmt"

	"github.com/chelnak/go-forge-yourself/forge"
)

func main() {
	client := forge.NewClient()

	ctx := context.Background()

	opts := &forge.ListReleasesOptions{
		Limit: 100,
		Owner: "puppetlabs",
	}

	response, err := client.Releases.ListReleases(ctx, opts)
	if err != nil {
		panic(err)
	}

	for _, release := range response.Results {
		fmt.Println(release.Slug)
	}
}

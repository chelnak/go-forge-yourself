package forge_test

import (
	"context"
	"testing"

	"github.com/chelnak/go-forge-yourself/forge"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestDeleteModuleSuccess(t *testing.T) {
	defer gock.Off()

	gock.New(forge.DefaultBaseURL).
		Delete("/modules/puppetlabs-stdlib").
		Reply(204)

	client := forge.NewClient()
	modules := client.Modules

	ctx := context.Background()
	err := modules.DeleteModule(ctx, "puppetlabs-stdlib", forge.DeleteModuleParams{})
	assert.NoError(t, err)
}

func TestDeleteModuleError(t *testing.T) {
	defer gock.Off()

	gock.New(forge.DefaultBaseURL).
		Delete("/modules/puppetlabs-stdlib").
		Reply(404).
		JSON(forge.Error{
			Message: "404 Module not found",
			Errors:  []string{"Module not found"},
		})

	client := forge.NewClient()
	modules := client.Modules

	ctx := context.Background()
	err := modules.DeleteModule(ctx, "puppetlabs-stdlib", forge.DeleteModuleParams{})
	assert.Error(t, err)
	assert.Equal(t, "404 Module not found", err.Error())
}

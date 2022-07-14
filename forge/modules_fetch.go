package forge

import (
	"context"
	"encoding/json"
)

// GetModuleOptions represents the options that can be passed to the GetModule method.
type GetModuleOptions struct {
	WithHTML      bool     `url:"with_html,omitempty"`
	IncludeFields []string `url:"include_fields,omitempty"`
	ExcludeFields []string `url:"exclude_fields,omitempty"`
}

// GetModule returns a single module from the forge API. The response can be controlled by passing in a GetModuleOptions struct.
// https://forgeapi.puppet.com/#operation/getModule
func (s *ModulesService) GetModule(ctx context.Context, slug string, opts GetModuleOptions) (*Module, error) {
	moduleURI := modulesEndpoint + "/" + slug

	req, err := s.client.NewRequest(ctx, "GET", moduleURI, nil, opts)
	if err != nil {
		return nil, err
	}

	res, err := s.client.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = checkResponseError(res)
	if err != nil {
		return nil, err
	}

	module := new(Module)
	if err = json.NewDecoder(res.Body).Decode(module); err != nil {
		return nil, err
	}

	return module, nil
}

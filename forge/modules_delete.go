package forge

import "context"

type DeleteModuleParams struct {
	Reason string `url:"reason,omitempty"`
}

// DeleteModule Perform a soft delete of a module, identified by the module's slug value.
// https://forgeapi.puppet.com/#operation/deleteModule
func (s *ModulesService) DeleteModule(ctx context.Context, slug string, params DeleteModuleParams) error {
	moduleURI := modulesEndpoint + "/" + slug

	req, err := s.client.NewRequest(ctx, "DELETE", moduleURI, nil, params)
	if err != nil {
		return err
	}

	res, err := s.client.client.Do(req)
	if err != nil {
		return err
	} else if res.StatusCode != 204 {
		if err = checkResponseError(res); err != nil {
			return err
		}
	}

	return nil
}

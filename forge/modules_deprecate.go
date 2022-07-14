package forge

import "context"

// Parameters for module deprecation
type DeprecateModuleParams struct {
	Reason          string `json:"reason"`           // Reason for deprecation
	ReplacementSlug string `json:"replacement_slug"` // Slug identifying a replacement Module. Accepts legacy (case-sensitive) module naming
}

// Action and params for patch operation.
type DeprecateModuleBody struct {
	Action string                `json:"action"`
	Params DeprecateModuleParams `json:"params"`
}

// Mark a module, identified by the module's slug value, as "deprecated".
// https://forgeapi.puppet.com/#tag/Module-Operations/operation/deprecateModule
func (s *ModulesService) DeprecateModule(ctx context.Context, slug string, params DeprecateModuleParams) error {
	moduleURI := modulesEndpoint + "/" + slug

	body := DeprecateModuleBody{
		Action: "delete",
		Params: params,
	}
	req, err := s.client.NewRequest(ctx, "PATCH", moduleURI, body, nil)
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

package forge

import (
	"context"
	"encoding/json"
)

// ListModulesOptions represents the options that can be passed to the ListModules method.
type ListModulesOptions struct {
	Limit                  int           `url:"limit,omitempty"`
	Offset                 int           `url:"offset,omitempty"`
	SortBy                 SortOption    `url:"sort_by,omitempty"`
	Tag                    string        `url:"tag,omitempty"`
	Owner                  string        `url:"owner,omitempty"`
	WithTasks              bool          `url:"with_tasks,omitempty"`
	WithPlans              bool          `url:"with_plans,omitempty"`
	WithPDK                bool          `url:"with_pdk,omitempty"`
	Premium                bool          `url:"premium,omitempty"`
	ExcludePremium         bool          `url:"exclude_premium,omitempty"`
	Endorsements           []Endorsement `url:"endorsements,omitempty"`
	OperatingSystem        string        `url:"operating_system,omitempty"`
	OperatingSystemRelease string        `url:"operating_system_release,omitempty"`
	PERequirement          string        `url:"pe_requirement,omitempty"`
	PuppetRequirement      string        `url:"puppet_requirement,omitempty"`
	WithMinimumScore       int           `url:"with_minimum_score,omitempty"`
	ModuleGroups           []ModuleGroup `url:"module_groups,omitempty"`
	ShowDeleted            bool          `url:"show_deleted,omitempty"`
	HideDeprecated         bool          `url:"hide_deprecated,omitempty"`
	OnlyLatest             bool          `url:"only_latest,omitempty"`
	Slugs                  []string      `url:"slugs,omitempty"`
	WithHTML               bool          `url:"with_html,omitempty"`
	IncludeFields          []string      `url:"include_fields,omitempty"`
	ExcludeFields          []string      `url:"exclude_fields,omitempty"`
	StartsWith             string        `url:"starts_with,omitempty"`
	WithReleaseSince       string        `url:"with_release_since,omitempty"`
}

// ListModulesResponse is the response from the ListModules method.
type ListModulesResponse struct {
	Pagination Pagination `json:"pagination"`
	Results    []Module   `json:"results"`
}

// ListModules returns a list of modules from the forge API. The response can be controlled
// by passing in a ListModulesOptions struct.
// https://forgeapi.puppet.com/#operation/getModules
func (s *ModulesService) ListModules(ctx context.Context, opts *ListModulesOptions) (*ListModulesResponse, error) {
	req, err := s.client.NewRequest(ctx, "GET", modulesEndpoint, nil, opts)
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

	response := new(ListModulesResponse)
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

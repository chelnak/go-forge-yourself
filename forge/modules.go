package forge

import (
	"context"
	"encoding/json"
)

const (
	modulesEndpoint = "modules"

	SortByRank          SortOption = "rank"
	SortByDownloads     SortOption = "downloads"
	SortByLatestRelease SortOption = "latest_release"

	EndorsementSupported Endorsement = "supported"
	EndorsementApproved  Endorsement = "approved"
	EndorsementPartner   Endorsement = "partner"

	ModuleGroupBase ModuleGroup = "base"
	ModuleGroupUtil ModuleGroup = "pe_only"
)

//ModulesService represents services for interactivng with modules
// related endpoints of the forge API.
type ModulesService service

// SortOption represents the desired order in which to return results.
type SortOption string

// Endorsement indicates whether a module is endorsed through the
// Supported, Approved or Partner Supported programs, or null if
// not endorsed.
type Endorsement string

// ModuleGroup indicates whether or not a module is licensed for use by
// Puppet Enterprise customers only, indicated by a value of pe_only,
// for all other modules this value will be base.
type ModuleGroup string

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

// GetModuleOptions represents the options that can be passed to the GetModule method.
type GetModuleOptions struct {
	WithHTML      bool     `url:"with_html,omitempty"`
	IncludeFields []string `url:"include_fields,omitempty"`
	ExcludeFields []string `url:"exclude_fields,omitempty"`
}

// ListModulesResponse is the response from the ListModules method.
type ListModulesResponse struct {
	Pagination Pagination `json:"pagination"`
	Results    []Module   `json:"results"`
}

// Pagination represents the pagination information that is included in a response
// from the forge API.
type Pagination struct {
	Limit   int    `json:"limit,omitempty"`
	Offset  int    `json:"offset,omitempty"`
	First   string `json:"first,omitempty"`
	Prev    string `json:"prev,omitempty"`
	Current string `json:"current,omitempty"`
	Next    string `json:"next,omitempty"`
	Total   int    `json:"total,omitempty"`
}

// Module represents a module entity from the forge API.
type Module struct {
	URI            string      `json:"uri,omitempty"`
	Slug           string      `json:"slug,omitempty"`
	Name           string      `json:"name,omitempty"`
	Downloads      int         `json:"downloads,omitempty"`
	CreatedAt      string      `json:"created_at,omitempty"`
	UpdatedAt      string      `json:"updated_at,omitempty"`
	DeprecatedAt   string      `json:"deprecated_at,omitempty"`
	DeprecatedFor  string      `json:"deprecated_for,omitempty"`
	SupersededBy   struct{}    `json:"superseded_by,omitempty"`
	Endorsement    Endorsement `json:"endorsement,omitempty"`
	ModuleGroup    ModuleGroup `json:"module_group,omitempty"`
	Premium        bool        `json:"premium,omitempty"`
	Owner          struct{}    `json:"owner,omitempty"`           // object?
	CurrentRelease struct{}    `json:"current_release,omitempty"` // object?
	Releases       []struct{}  `json:"releases,omitempty"`        // array of object?
	FeedbackScore  int         `json:"feedback_score,omitempty"`
	IssuesUrl      string      `json:"issues_url,omitempty"`
}

// HasNext returns true if the there are more pages to return.
func (r *ListModulesResponse) HasNext() bool {
	return r.Pagination.Next != ""
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

	listModulesReponse := new(ListModulesResponse)
	if err = json.NewDecoder(res.Body).Decode(listModulesReponse); err != nil {
		return nil, err
	}

	return listModulesReponse, nil
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

	module := new(Module)
	if err = json.NewDecoder(res.Body).Decode(module); err != nil {
		return nil, err
	}

	return module, nil
}

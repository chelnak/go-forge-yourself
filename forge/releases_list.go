package forge

import (
	"context"
	"encoding/json"
)

// ListReleasesOptions represents the options that can be passed to the ListReleases method.
type ListReleasesOptions struct {
	Limit                  int      `url:"limit,omitempty"`
	Offset                 int      `url:"offset,omitempty"`
	SortBy                 string   `url:"sort_by,omitempty"`
	Module                 string   `url:"module,omitempty"`
	Owner                  string   `url:"owner,omitempty"`
	WithPDK                bool     `url:"with_pdk,omitempty"`
	OperatingSystem        string   `url:"operating_system,omitempty"`
	OperatingSystemRelease string   `url:"operating_system_release,omitempty"`
	PERequirement          string   `url:"pe_requirement,omitempty"`
	PuppetRequirement      string   `url:"puppet_requirement,omitempty"`
	ModuleGroups           []string `url:"module_groups,omitempty"`
	ShowDeleted            bool     `url:"show_deleted,omitempty"`
	HideDeprecated         bool     `url:"hide_deprecated,omitempty"`
	WithHTML               bool     `url:"with_html,omitempty"`
	IncludeFields          []string `url:"include_fields,omitempty"`
	ExcludeFields          []string `url:"exclude_fields,omitempty"`
}

// ListReleasesResponse represents the response from the ListReleases method.
type ListReleasesResponse struct {
	Pagination Pagination `json:"pagination"`
	Results    []Release  `json:"results"`
}

// ListReleases returns a list of module releases meeting the specified search
// criteria and filters. Results are paginated. All of the parameters are optional.
func (s *ReleasesService) ListReleases(ctx context.Context, opts *ListReleasesOptions) (*ListReleasesResponse, error) {
	req, err := s.client.NewRequest(ctx, "GET", releasesEndpoint, nil, opts)
	if err != nil {
		return nil, err
	}

	res, err := s.client.client.Do(req)
	if err != nil {
		return nil, err
	}

	response := new(ListReleasesResponse)
	if err = json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}

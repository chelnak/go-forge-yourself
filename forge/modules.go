package forge

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

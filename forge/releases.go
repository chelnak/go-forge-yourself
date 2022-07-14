package forge

const (
	releasesEndpoint = "releases"
)

// ReleasesService represents services for interactivng with releases
// related endpoints of the forge API.
type ReleasesService service

// Release represents a module release entity from the forge API.
type Release struct {
	URI             string     `json:"uri,omitempty"`
	Slug            string     `json:"slug,omitempty"`
	Module          Module     `json:"module,omitempty"`
	Version         string     `json:"version,omitempty"`
	Metadata        struct{}   `json:"metadata,omitempty"`
	Tags            []string   `json:"tags,omitempty"`
	PDK             bool       `json:"pdk,omitempty"`
	ValidationScore int        `json:"validation_score,omitempty"`
	FileURI         string     `json:"file_uri,omitempty"`
	FileSize        int        `json:"file_size,omitempty"`
	FileMD5         string     `json:"file_md5,omitempty"`
	FileSHA256      string     `json:"file_sha256,omitempty"`
	Downloads       int        `json:"downloads,omitempty"`
	Readme          string     `json:"readme,omitempty"`
	ChangeLog       string     `json:"changelog,omitempty"`
	License         string     `json:"license,omitempty"`
	References      string     `json:"references,omitempty"`
	PECompatibility []string   `json:"pe_compatibility,omitempty"`
	Tasks           []struct{} `json:"tasks,omitempty"`
	Plans           []struct{} `json:"plans,omitempty"`
	CreatedAt       string     `json:"created_at,omitempty"`
	UpdatedAt       string     `json:"updated_at,omitempty"`
	DeletedFor      string     `json:"deleted_for,omitempty"`
}

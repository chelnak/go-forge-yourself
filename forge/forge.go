package forge

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

// Constants
const (
	defaultBaseURL   = "https://forgeapi.puppet.com/v3/"
	defaultUserAgent = "go-forge-yourself/0.0.0"
)

// Common service structure
type service struct {
	client *Client
}

// Error represents a the basic error response returned by forge API requests.
type Error struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
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

// HasNext returns true if the there are more pages to return.
func (p *Pagination) HasNext() bool {
	return p.Next != ""
}

// Client represents a forge http client instance.
type Client struct {
	BaseURL   *url.URL
	UserAgent string
	apiKey    string
	client    *http.Client
	common    service

	Modules  *ModulesService
	Releases *ReleasesService
}

// ClientOption represents an option that can be passed to NewClient.
type ClientOption func(*Client)

// NewClient returns a new Forge API Client. It can be configured with the
// following options: BaseUrl, UserAgent, HTTPClient.
func NewClient(options ...ClientOption) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:    &http.Client{},
		BaseURL:   baseURL,
		apiKey:    "",
		UserAgent: defaultUserAgent,
	}

	for _, option := range options {
		option(c)
	}

	c.common.client = c
	c.Modules = (*ModulesService)(&c.common)
	c.Releases = (*ReleasesService)(&c.common)

	return c
}

// WithBaseUrl sets the base URL for the client. This defaults to the value
// of defaultBaseUrl.
// https://forgeapi.puppet.com/#section/Hostname-Configuration
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		u, _ := url.Parse(baseURL)
		c.BaseURL = u
	}
}

// WithHTTPClient sets the HTTP client to use for the client. This defaults to
// a basic http.Client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.client = httpClient
	}
}

// WithUserAgent sets the User-Agent header. This defaults to the value of
// defaultUserAgent. If you want to specify your own user agent, use
// WithUserAgent(userAgent string).
// https://forgeapi.puppet.com/#section/User-Agent-Required
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) {
		c.UserAgent = userAgent
	}
}

// WithAuthentication sets the Authentication header on the http request.
// https://forgeapi.puppet.com/#section/Authentication
func WithAuthentication(apiKey string) ClientOption {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash. If specified, the
// value pointed to by body is JSON encoded and included in as the request body.
func (c *Client) NewRequest(ctx context.Context, method string, urlStr string, body interface{}, opts interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	var err error
	if opts != nil {
		urlStr, err = addOptions(urlStr, opts)
		if err != nil {
			return nil, err
		}
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.apiKey != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// checkResponseError checks if response is an error.
func checkResponseError(res *http.Response) error {
	if res.StatusCode < 300 {
		return nil
	}

	errRes := new(Error)
	err := json.NewDecoder(res.Body).Decode(&errRes)

	if err != nil {
		return err
	} else if errRes.Message != "" {
		return errors.New(errRes.Message)
	}

	return nil
}

// addOptions takes a URL and a list of query options and returns a new URL
// with the options encoded as a query string.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

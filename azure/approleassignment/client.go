package approleassignment

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// NewClient returns a new Azure API client.
// Needs tenantId, clientId and clientSecret as input to generate an authenticated httpClient
// Uses golang.org/x/oauth2/clientcredentials pkg and returns a new configured AzureClient
func NewClient(tenantID, clientID, clientSecret string) *AzureClient {
	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     fmt.Sprintf(GRAPH_TOKEN_URL, tenantID),
		Scopes:       []string{GRAPH_DEFAULT_SCOPE},
	}

	ctx := context.Background()

	client := config.Client(ctx)

	url, _ := url.Parse(GRAPH_API_BASE_URL)

	return &AzureClient{tenantID: tenantID, httpClient: client, baseURL: url}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *AzureClient) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
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

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded , or returned as an error if an API error has occurred.
// The provided ctx must be non-nil, if it is nil an error is returned. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *AzureClient) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return nil, err
		}
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range or equal to 202 Accepted.
// API error responses are expected to have response
// body, and a JSON response body that maps to oDataResponse.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	odataResponse := &odataResponse{Resp: r}
	data, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		return fmt.Errorf("Error Reading Response Body : %s", readErr)
	}
	if readErr == nil && data != nil {
		if err := json.Unmarshal(data, &odataResponse); err != nil {
			return err
		}
	}
	// Re-populate error response body
	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	return odataResponse
}

package base

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	// BaseURL Goong API base URL
	BaseURL = "https://rsapi.goong.io"

	statusRateLimitExceeded = 429
)

// Base Goong API base
type Base struct {
	apiKey string
	debug  bool
}

// NewBase Create a new API base instance
func NewBase(apiKey string) (*Base, error) {
	if apiKey == "" {
		return nil, errors.New("Goong api_key not found")
	}

	b := &Base{}

	b.apiKey = apiKey

	return b, nil
}

// SetDebug enables debug output for API calls
func (b *Base) SetDebug(debug bool) {
	b.debug = true
}

type GoongApiError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// QueryRequest make a get with the provided query string and return the response if successful
func (b *Base) QueryRequest(query string, v *url.Values) (*http.Response, error) {
	// Add api_key to args
	v.Set("api_key", b.apiKey)

	// Generate URL
	url := fmt.Sprintf("%s/%s", BaseURL, query)

	if b.debug {
		fmt.Printf("URL: %s\n", url)
	}

	// Create request object
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = v.Encode()

	// Create client instance
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if b.debug {
		data, _ := httputil.DumpRequest(request, true)
		fmt.Printf("Request: %s", string(data))
		data, _ = httputil.DumpResponse(resp, false)
		fmt.Printf("Response: %s", string(data))
	}

	if resp.StatusCode == statusRateLimitExceeded {
		return nil, ErrorAPILimitExceeded
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, ErrorAPIUnauthorized
	}

	return resp, nil
}

// QueryBase Query the Goong API and fill the provided instance with the returned JSON
// TODO: Rename this
func (b *Base) QueryBase(query string, v *url.Values, inst interface{}) error {
	// Make request
	resp, err := b.QueryRequest(query, v)
	if err != nil && (resp == nil || resp.StatusCode != http.StatusBadRequest) {
		return err
	}
	defer resp.Body.Close()

	// Read body into buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Handle bad requests with messages
	if resp.StatusCode != http.StatusOK {
		apiMessage := GoongApiError{}
		messageErr := json.Unmarshal(body, &apiMessage)
		if messageErr == nil {
			return fmt.Errorf("api error: %s", apiMessage.Error.Message)
		}
		return fmt.Errorf("Bad Request (400) - no message")
	}

	// Attempt to decode body into inst type
	err = json.Unmarshal(body, &inst)
	if err != nil {
		return err
	}

	return nil
}

// Query the Goong API
// TODO: Depreciate this
func (b *Base) Query(api, mode string, v *url.Values, inst interface{}) error {
	// Generate URL
	queryString := fmt.Sprintf("%s/%s", api, mode)
	return b.QueryBase(queryString, v, inst)
}

package revenuecat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client makes authorized calls to the RevenueCat API.
type Client struct {
	apiKey string
	apiURL string
	http   doer
}

type doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// New returns a new *Client for the provided API key.
// For more information on authentication, see https://docs.revenuecat.com/docs/authentication.
func New(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		apiURL: "https://api.revenuecat.com/v1/",
		http: &http.Client{
			// Set a long timeout here since calls to Apple are probably invloved.
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) call(method, path string, reqBody interface{}, platform string, respBody interface{}) error {
	var reqBodyJSON io.Reader
	if reqBody != nil {
		js, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("error marshaling request body: %v", err)
		}
		reqBodyJSON = bytes.NewBuffer(js)
	}
	req, err := http.NewRequest(method, c.apiURL+path, reqBodyJSON)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.Header.Add("Content-Type", "application/json")

	if platform == "" {
		platform = "ios"
	}
	req.Header.Add("X-Platform", platform)

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		var errResp Error
		err = json.NewDecoder(resp.Body).Decode(&errResp)
		if err != nil {
			return err
		}
		return errResp
	}
	if respBody == nil {
		// Expecting an empty body.
		return nil
	}
	err = json.NewDecoder(resp.Body).Decode(respBody)
	if err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}
	return nil
}

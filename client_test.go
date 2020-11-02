package revenuecat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type mockClient struct {
	request *http.Request
	doer    func(req *http.Request) (*http.Response, error)
}

func newMockClient(t *testing.T, statusCode int, body interface{}, returnErr error) *mockClient {
	t.Helper()

	c := &mockClient{}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("unable to marshal response body: %v", err)
	}

	c.doer = func(req *http.Request) (*http.Response, error) {
		resp := &http.Response{
			StatusCode: statusCode,
			Body:       ioutil.NopCloser(bytes.NewReader(bodyBytes)),
		}
		return resp, returnErr
	}

	return c
}

func (c *mockClient) Do(req *http.Request) (*http.Response, error) {
	c.request = req
	return c.doer(req)
}

func (c *mockClient) expectPath(t *testing.T, expected string) {
	t.Helper()
	if path := c.request.URL.Path; path != expected {
		t.Errorf("expected path: %q, got path: %q", expected, path)
	}
}

func (c *mockClient) expectMethod(t *testing.T, expected string) {
	t.Helper()
	if method := c.request.Method; method != expected {
		t.Errorf("expected method: %q, got method: %q", expected, method)
	}
}

func (c *mockClient) expectXPlatform(t *testing.T, expected string) {
	t.Helper()
	if platform := c.request.Header.Get("x-platform"); platform != expected {
		t.Errorf("expected x-platform: %q, got x-platform: %q", expected, platform)
	}
}

func (c *mockClient) expectBody(t *testing.T, expected string) {
	t.Helper()
	bodyBytes, err := ioutil.ReadAll(c.request.Body)
	if err != nil {
		t.Fatalf("error reading request body: %v", err)
	}

	body := string(bodyBytes)

	if body != expected {
		t.Errorf("expected body: %q\n, got body: %q", expected, body)
	}
}

func staticTime(t *testing.T, timeStr string) time.Time {
	t.Helper()
	if timeStr == "" {
		return time.Time{}
	}
	val, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		t.Fatalf("unable to create time from %q, %v", timeStr, err)
	}
	return val
}

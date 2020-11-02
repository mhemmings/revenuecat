package revenuecat

import (
	"testing"
)

func TestOverrideOffering(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	_, err := rc.OverrideOffering("123", "testUUID")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/subscribers/123/offerings/testUUID/override")
}

func TestDeleteOfferingOverride(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	_, err := rc.DeleteOfferingOverride("123")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "DELETE")
	cl.expectPath(t, "/v1/subscribers/123/offerings/override")
}

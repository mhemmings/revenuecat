package revenuecat

import (
	"testing"
)

func TestAddUserAttribution(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	err := rc.AddUserAttribution("123", Facebook, AttributionData{IDFA: "test"})
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/subscribers/123/attribution")
	cl.expectBody(t, `{"data":{"rc_idfa":"test"},"network":5}`)
}

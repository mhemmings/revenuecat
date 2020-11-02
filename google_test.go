package revenuecat

import (
	"testing"
)

func TestRefundGoogleSubscription(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	_, err := rc.RefundGoogleSubscription("123", "sub")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/subscribers/123/subscriptions/sub/revoke")
}

func TestDeferGoogleSubscription(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	_, err := rc.DeferGoogleSubscription("123", "sub", staticTime(t, "2020-01-15 23:54:17"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/subscribers/123/subscriptions/sub/defer")
	cl.expectBody(t, `{"expiry_time_ms":1579132457000}`)
}

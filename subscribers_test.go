package revenuecat

import (
	"encoding/json"
	"testing"
	"time"
)

func TestGetSubscriber(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	_, err := rc.GetSubscriber("123")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "GET")
	cl.expectPath(t, "/v1/subscribers/123")
}

func TestGetSubscriberSandbox(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey", WithSandboxEnabled(true))
	rc.http = cl

	_, err := rc.GetSubscriber("123")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "GET")
	cl.expectPath(t, "/v1/subscribers/123")
	cl.expectXIsSandbox(t, "true")
}

func TestGetSubscriberSandboxDisabled(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey", WithSandboxEnabled(false))
	rc.http = cl

	_, err := rc.GetSubscriber("123")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "GET")
	cl.expectPath(t, "/v1/subscribers/123")
	cl.expectXIsSandbox(t, "")
}

func TestGetSubscriberWithPlatform(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	_, err := rc.GetSubscriberWithPlatform("123", "ios")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "GET")
	cl.expectPath(t, "/v1/subscribers/123")
	cl.expectXPlatform(t, "ios")
}

func TestUpdateSubscriberAttributes(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	attrs := map[string]SubscriberAttribute{
		"foo": {Value: "bar"},
		"x": {
			Value:     "y",
			UpdatedAt: staticTime(t, "2020-01-15 23:54:17"),
		},
	}
	err := rc.UpdateSubscriberAttributes("123", attrs)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/subscribers/123/attributes")
	cl.expectBody(t, `{"attributes":{"foo":{"value":"bar"},"x":{"value":"y","updated_at_ms":1579132457000}}}`)
}

func TestDeleteSubscriber(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	err := rc.DeleteSubscriber("123")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "DELETE")
	cl.expectPath(t, "/v1/subscribers/123")
}

func TestSubscriberAttributeMarshalJSON(t *testing.T) {
	attr := SubscriberAttribute{
		Value:     "foo",
		UpdatedAt: time.Unix(0, 1554130937000000000),
	}

	b, err := json.Marshal(attr)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	expected := `{"value":"foo","updated_at_ms":1554130937000}`
	if string(b) != expected {
		t.Errorf("expected: %s, actual: %s", expected, string(b))
	}
}

func TestSubscriberAttributeUnmarshalJSON(t *testing.T) {
	attr := `{"value":"foo","updated_at_ms":1554130937000}`
	var res SubscriberAttribute
	err := json.Unmarshal([]byte(attr), &res)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	expected := SubscriberAttribute{
		Value:     "foo",
		UpdatedAt: time.Unix(0, 1554130937000000000),
	}

	if res != expected {
		t.Errorf("expected: %q, actual: %q", expected, res)
	}
}

func TestSubscriberIsEntitledTo(t *testing.T) {
	tests := []struct {
		name        string
		sub         map[string]Entitlement
		entitlement string
		expected    bool
	}{{
		name:        "nil",
		sub:         nil,
		entitlement: "test",
		expected:    false,
	}, {
		name:        "empty",
		sub:         make(map[string]Entitlement),
		entitlement: "test",
		expected:    false,
	}, {
		name: "missing",
		sub: map[string]Entitlement{
			"foo": {},
		},
		entitlement: "test",
		expected:    false,
	}, {
		name: "expired",
		sub: map[string]Entitlement{
			"test": {
				ExpiresDate: time.Now().Add(-time.Hour),
			},
		},
		entitlement: "test",
		expected:    false,
	}, {
		name: "subscribed",
		sub: map[string]Entitlement{
			"test": {
				ExpiresDate: time.Now().Add(time.Hour),
			},
		},
		entitlement: "test",
		expected:    true,
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := Subscriber{
				Entitlements: test.sub,
			}
			if res := s.IsEntitledTo(test.entitlement); res != test.expected {
				t.Errorf("got: %v, expected %v", res, test.expected)
			}
		})
	}
}

package revenuecat

import (
	"testing"
)

func TestCreatePurchaseWithoutOpts(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	_, err := rc.CreatePurchase("123", "testreceipt", nil)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/receipts")
	cl.expectXPlatform(t, "ios")
	cl.expectBody(t, `{"app_user_id":"123","fetch_token":"testreceipt"}`)
}

func TestCreatePurchaseWithOpts(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	opt := &CreatePurchaseOptions{
		Platform: "ios",

		ProductID: "product_sku",
		IsRestore: true,
	}

	_, err := rc.CreatePurchase("123", "testreceipt", opt)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/receipts")
	cl.expectXPlatform(t, "ios")
	cl.expectBody(t, `{"app_user_id":"123","fetch_token":"testreceipt","product_id":"product_sku","is_restore":true}`)
}

func TestCreatePurchaseWithStripe(t *testing.T) {
	cl := newMockClient(t, 200, nil, nil)
	rc := New("apikey")
	rc.http = cl

	opt := &CreatePurchaseOptions{
		Platform: string(StripeStore),

		ProductID: "product_sku",
		IsRestore: true,
	}

	_, err := rc.CreatePurchase("123", "testreceipt", opt)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	cl.expectMethod(t, "POST")
	cl.expectPath(t, "/v1/receipts")
	cl.expectXPlatform(t, "stripe")
	cl.expectBody(t, `{"app_user_id":"123","fetch_token":"testreceipt","product_id":"product_sku","is_restore":true}`)
}

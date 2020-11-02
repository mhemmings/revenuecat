package revenuecat

import "time"

// RefundGoogleSubscription immediately revokes access to a Google Subscription and issues a refund for the last purchase.
// https://docs.revenuecat.com/reference#revoke-a-google-subscription
func (c *Client) RefundGoogleSubscription(userID string, id string) (Subscriber, error) {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}

	err := c.call("POST", "subscribers/"+userID+"/subscriptions/"+id+"/revoke", nil, "", &resp)
	return resp.Subscriber, err
}

// DeferGoogleSubscription defers the purchase of a Google Subscription to a later date.
// https://docs.revenuecat.com/reference#defer-a-google-subscription
func (c *Client) DeferGoogleSubscription(userID string, id string, nextExpiry time.Time) (Subscriber, error) {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}

	req := struct {
		ExpiryTime int64 `json:"expiry_time_ms,omitempty"`
	}{
		ExpiryTime: toMilliseconds(nextExpiry),
	}

	err := c.call("POST", "subscribers/"+userID+"/subscriptions/"+id+"/defer", req, "", &resp)
	return resp.Subscriber, err
}

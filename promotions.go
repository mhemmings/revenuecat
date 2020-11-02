package revenuecat

import "time"

// GrantEntitlement grants a user a promotional entitlement.
// https://docs.revenuecat.com/reference#grant-a-promotional-entitlement
func (c *Client) GrantEntitlement(userID string, id string, duration Duration, startTime time.Time) (Subscriber, error) {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}

	req := struct {
		Duration  Duration `json:"duration"`
		StartTime int64    `json:"start_time_ms,omitempty"`
	}{
		Duration: duration,
	}

	if !startTime.IsZero() {
		req.StartTime = toMilliseconds(startTime)
	}

	err := c.call("POST", "subscribers/"+userID+"/entitlements/"+id+"/promotional", req, "", &resp)
	return resp.Subscriber, err
}

// RevokeEntitlement revokes all promotional entitlements for a given entitlement identifier and app user ID.
// https://docs.revenuecat.com/reference#revoke-promotional-entitlements
func (c *Client) RevokeEntitlement(userID string, id string) (Subscriber, error) {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}

	err := c.call("POST", "subscribers/"+userID+"/entitlements/"+id+"/revoke_promotionals", nil, "", &resp)
	return resp.Subscriber, err
}

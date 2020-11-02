package revenuecat

// OverrideOffering overrides the current Offering for a specific user.
// https://docs.revenuecat.com/reference#override-offering
func (c *Client) OverrideOffering(userID string, offeringUUID string) (Subscriber, error) {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}
	err := c.call("POST", "subscribers/"+userID+"/offerings/"+offeringUUID+"/override", nil, "", &resp)
	return resp.Subscriber, err
}

// DeleteOfferingOverride reset the offering overrides back to the current offering for a specific user.
// https://docs.revenuecat.com/reference#delete-offering-override
func (c *Client) DeleteOfferingOverride(userID string) (Subscriber, error) {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}
	err := c.call("DELETE", "subscribers/"+userID+"/offerings/override", nil, "", &resp)
	return resp.Subscriber, err
}

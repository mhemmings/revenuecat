package revenuecat

// Network represents a predefined attribution channel.
type Network int

// https://docs.revenuecat.com/reference#attribution-source-network-codes
const (
	AppleSearchAds Network = iota
	Adjust
	AppsFlyer
	Branch
	Tenjin
	Facebook
)

// AttributionData holds the identifier value for either the App Store or Play Services.
type AttributionData struct {
	IDFA           string `json:"rc_idfa,omitempty"`
	PlayServicesID string `json:"rc_gps_adid,omitempty"`
}

// AddUserAttribution attaches attribution data to a subscriber from specific supported networks.
// https://docs.revenuecat.com/reference#subscribersattribution
func (c *Client) AddUserAttribution(userID string, network Network, data AttributionData) error {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}

	req := struct {
		Data    AttributionData `json:"data"`
		Network Network         `json:"network"`
	}{
		Data:    data,
		Network: network,
	}

	return c.call("POST", "subscribers/"+userID+"/attribution", req, "", &resp)
}

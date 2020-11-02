package revenuecat

// CreatePurchaseOptions holds the optional values for creating a purchase.
// https://docs.revenuecat.com/reference#receipts
type CreatePurchaseOptions struct {
	Platform string `json:"-"`

	ProductID         string                         `json:"product_id,omitempty"`
	Price             float32                        `json:"price,omitempty"`
	Currency          string                         `json:"currency,omitempty"`
	PaymentMode       string                         `json:"payment_mode,omitempty"`
	IntroductoryPrice float32                        `json:"introductory_price,omitempty"`
	IsRestore         bool                           `json:"is_restore,omitempty"`
	Attributes        map[string]SubscriberAttribute `json:"attributes,omitempty"`
}

// CreatePurchase records a purchase for a user from iOS, Android, or Stripe and will create a user if they don't already exist.
// https://docs.revenuecat.com/reference#receipts
func (c *Client) CreatePurchase(userID string, receipt string, opt *CreatePurchaseOptions) (Subscriber, error) {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}

	req := struct {
		AppUserID  string `json:"app_user_id"`
		FetchToken string `json:"fetch_token"`
		*CreatePurchaseOptions
	}{
		AppUserID:             userID,
		FetchToken:            receipt,
		CreatePurchaseOptions: opt,
	}

	var platform string
	if opt != nil {
		platform = opt.Platform
	}

	err := c.call("POST", "receipts", req, platform, &resp)
	return resp.Subscriber, err
}

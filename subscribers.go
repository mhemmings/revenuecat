package revenuecat

import (
	"encoding/json"
	"time"
)

// Subscriber holds a subscriber returned by the RevenueCat API.
type Subscriber struct {
	OriginalAppUserID          string                         `json:"original_app_user_id"`
	OriginalApplicationVersion *string                        `json:"original_application_version"`
	FirstSeen                  time.Time                      `json:"first_seen"`
	LastSeen                   time.Time                      `json:"last_seen"`
	Entitlements               map[string]Entitlement         `json:"entitlements"`
	Subscriptions              map[string]Subscription        `json:"subscriptions"`
	NonSubscriptions           map[string][]NonSubscription   `json:"non_subscriptions"`
	SubscriberAttributes       map[string]SubscriberAttribute `json:"subscriber_attributes"`
}

// https://docs.revenuecat.com/reference#the-entitlement-object
type Entitlement struct {
	ExpiresDate            time.Time  `json:"expires_date"`
	GracePeriodExpiresDate *time.Time `json:"grace_period_expires_date"`
	PurchaseDate           time.Time  `json:"purchase_date"`
	ProductIdentifier      string     `json:"product_identifier"`
	ProductPlanIdentifier  string     `json:"product_plan_identifier"`
}

// https://docs.revenuecat.com/reference#the-subscription-object
type Subscription struct {
	ExpiresDate             *time.Time    `json:"expires_date"`
	PurchaseDate            time.Time     `json:"purchase_date"`
	OriginalPurchaseDate    time.Time     `json:"original_purchase_date"`
	PeriodType              PeriodType    `json:"period_type"`
	Store                   Store         `json:"store"`
	IsSandbox               bool          `json:"is_sandbox"`
	UnsubscribeDetectedAt   *time.Time    `json:"unsubscribe_detected_at"`
	BillingIssuesDetectedAt *time.Time    `json:"billing_issues_detected_at"`
	AutoResumeDate          *time.Time    `json:"auto_resume_date"`
	GracePeriodExpiresDate  *time.Time    `json:"grace_period_expires_date"`
	RefundedAt              *time.Time    `json:"refunded_at"`
	OwnershipType           OwnershipType `json:"ownership_type"`
	StoreTransactionID      string        `json:"store_transaction_id"`
	ProductPlanIdentifier   string        `json:"product_plan_identifier"`
}

// https://docs.revenuecat.com/reference#section-the-non-subscription-object
type NonSubscription struct {
	ID           string    `json:"id"`
	PurchaseDate time.Time `json:"purchase_date"`
	Store        Store     `json:"store"`
	IsSandbox    bool      `json:"is_sandbox"`
}

// https://docs.revenuecat.com/reference#section-the-subscriber-attribute-object
type SubscriberAttribute struct {
	Value     string
	UpdatedAt time.Time
}

// PeriodType holds the predefined values for a subscription period.
type PeriodType string

// https://docs.revenuecat.com/reference#the-subscription-object
const (
	NormalPeriodType PeriodType = "normal"
	TrialPeriodType  PeriodType = "trial"
	IntroPeriodType  PeriodType = "intro"
)

// OwnershipType holds the predefined values for a subscription ownership type.
type OwnershipType string

const (
	PurchasedOwnershipType    OwnershipType = "PURCHASED"
	FamilySharedOwnershipType OwnershipType = "FAMILY_SHARED"
)

// Store holds the predefined values for a store.
type Store string

// https://docs.revenuecat.com/reference#the-subscription-object
const (
	AppStore         Store = "app_store"
	MacAppStore      Store = "mac_app_store"
	PlayStore        Store = "play_store"
	StripeStore      Store = "stripe"
	PromotionalStore Store = "promotional"
)

// IsEntitledTo returns true if the Subscriber has the given entitlement.
func (s Subscriber) IsEntitledTo(entitlement string) bool {
	e, ok := s.Entitlements[entitlement]
	if !ok {
		return false
	}
	return !e.ExpiresDate.Before(time.Now())
}

// GetSubscriber gets the latest subscriber info or creates one if it doesn't exist.
// https://docs.revenuecat.com/reference#subscribers
func (c *Client) GetSubscriber(userID string) (Subscriber, error) {
	return c.GetSubscriberWithPlatform(userID, "")
}

// GetSubscriberWithPlatform gets the latest subscriber info or creates one if it doesn't exist, updating the subscriber record's last_seen
// value for the platform provided.
// https://docs.revenuecat.com/reference#subscribers
func (c *Client) GetSubscriberWithPlatform(userID string, platform string) (Subscriber, error) {
	var resp struct {
		Subscriber Subscriber `json:"subscriber"`
	}
	err := c.call("GET", "subscribers/"+userID, nil, platform, &resp)
	return resp.Subscriber, err
}

// UpdateSubscriberAttributes updates subscriber attributes for a user.
// https://docs.revenuecat.com/reference#update-subscriber-attributes
func (c *Client) UpdateSubscriberAttributes(userID string, attributes map[string]SubscriberAttribute) error {
	req := struct {
		Attributes map[string]SubscriberAttribute `json:"attributes"`
	}{
		Attributes: attributes,
	}
	return c.call("POST", "subscribers/"+userID+"/attributes", req, "", nil)
}

// DeleteSubscriber permanently deletes a subscriber.
// https://docs.revenuecat.com/reference#subscribersapp_user_id
func (c *Client) DeleteSubscriber(userID string) error {
	return c.call("DELETE", "subscribers/"+userID, nil, "", nil)
}

func (attr SubscriberAttribute) MarshalJSON() ([]byte, error) {
	var updatedAt int64
	if !attr.UpdatedAt.IsZero() {
		updatedAt = toMilliseconds(attr.UpdatedAt)
	}
	return json.Marshal(&struct {
		Value     string `json:"value"`
		UpdatedAt int64  `json:"updated_at_ms,omitempty"`
	}{
		Value:     attr.Value,
		UpdatedAt: updatedAt,
	})
}

func (attr *SubscriberAttribute) UnmarshalJSON(data []byte) error {
	var jsonAttr struct {
		Value     string `json:"value"`
		UpdatedAt int64  `json:"updated_at_ms,omitempty"`
	}
	if err := json.Unmarshal(data, &jsonAttr); err != nil {
		return err
	}
	attr.Value = jsonAttr.Value
	if jsonAttr.UpdatedAt > 0 {
		attr.UpdatedAt = fromMilliseconds(jsonAttr.UpdatedAt)
	}
	return nil
}

# RevenueCat

[![PkgGoDev](https://pkg.go.dev/badge/github.com/mhemmings/revenuecat)](https://pkg.go.dev/github.com/mhemmings/revenuecat)
[![Test](https://github.com/mhemmings/revenuecat/workflows/Test/badge.svg?branch=master)](https://github.com/mhemmings/revenuecat/actions?query=workflow%3ATest)

Go package for interacting with the [RevenueCat API](https://docs.revenuecat.com/reference).

## Usage

### Example

```go
package main

import (
	"fmt"

	"github.com/mhemmings/revenuecat"
)

func main() {
	rc := revenuecat.New("apikey")
	sub, _ := rc.GetSubscriber("123")
	entitled := sub.IsEntitledTo("premium")

	fmt.Println("user entitled: %t", entitled)
}
```

### Documentation

For full documentation, see [pkg.go.dev/github.com/mhemmings/revenuecat](https://pkg.go.dev/github.com/mhemmings/revenuecat)


#### func (*Client) AddUserAttribution

```go
func (c *Client) AddUserAttribution(userID string, network Network, data AttributionData) error
```
AddUserAttribution attaches attribution data to a subscriber from specific
supported networks. https://docs.revenuecat.com/reference#subscribersattribution

#### func (*Client) CreatePurchase

```go
func (c *Client) CreatePurchase(userID string, receipt string, opt *CreatePurchaseOptions) (Subscriber, error)
```
CreatePurchase records a purchase for a user from iOS, Android, or Stripe and
will create a user if they don't already exist.
https://docs.revenuecat.com/reference#receipts

#### func (*Client) DeferGoogleSubscription

```go
func (c *Client) DeferGoogleSubscription(userID string, id string, nextExpiry time.Time) (Subscriber, error)
```
DeferGoogleSubscription defers the purchase of a Google Subscription to a later
date. https://docs.revenuecat.com/reference#defer-a-google-subscription

#### func (*Client) DeleteOfferingOverride

```go
func (c *Client) DeleteOfferingOverride(userID string) (Subscriber, error)
```
DeleteOfferingOverride reset the offering overrides back to the current offering
for a specific user.
https://docs.revenuecat.com/reference#delete-offering-override

#### func (*Client) DeleteSubscriber

```go
func (c *Client) DeleteSubscriber(userID string) error
```
DeleteSubscriber permanently deletes a subscriber.
https://docs.revenuecat.com/reference#subscribersapp_user_id

#### func (*Client) GetSubscriber

```go
func (c *Client) GetSubscriber(userID string) (Subscriber, error)
```
GetSubscriber gets the latest subscriber info or creates one if it doesn't
exist. https://docs.revenuecat.com/reference#subscribers

#### func (*Client) GetSubscriberWithPlatform

```go
func (c *Client) GetSubscriberWithPlatform(userID string, platform string) (Subscriber, error)
```
GetSubscriberWithPlatform gets the latest subscriber info or creates one if it
doesn't exist, updating the subscriber record's last_seen value for the platform
provided. https://docs.revenuecat.com/reference#subscribers

#### func (*Client) GrantEntitlement

```go
func (c *Client) GrantEntitlement(userID string, id string, duration Duration, startTime time.Time) (Subscriber, error)
```
GrantEntitlement grants a user a promotional entitlement.
https://docs.revenuecat.com/reference#grant-a-promotional-entitlement

#### func (*Client) OverrideOffering

```go
func (c *Client) OverrideOffering(userID string, offeringUUID string) (Subscriber, error)
```
OverrideOffering overrides the current Offering for a specific user.
https://docs.revenuecat.com/reference#override-offering

#### func (*Client) RefundGoogleSubscription

```go
func (c *Client) RefundGoogleSubscription(userID string, id string) (Subscriber, error)
```
RefundGoogleSubscription immediately revokes access to a Google Subscription and
issues a refund for the last purchase.
https://docs.revenuecat.com/reference#revoke-a-google-subscription

#### func (*Client) RevokeEntitlement

```go
func (c *Client) RevokeEntitlement(userID string, id string) (Subscriber, error)
```
RevokeEntitlement revokes all promotional entitlements for a given entitlement
identifier and app user ID.
https://docs.revenuecat.com/reference#revoke-promotional-entitlements

#### func (*Client) UpdateSubscriberAttributes

```go
func (c *Client) UpdateSubscriberAttributes(userID string, attributes map[string]SubscriberAttribute) error
```
UpdateSubscriberAttributes updates subscriber attributes for a user.
https://docs.revenuecat.com/reference#update-subscriber-attributes


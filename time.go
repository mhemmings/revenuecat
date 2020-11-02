package revenuecat

import "time"

// Duration holds a predefined entitlement duration.
type Duration string

// https://docs.revenuecat.com/reference#duration-values
const (
	Daily      Duration = "daily"
	Weekly     Duration = "weekly"
	Monthly    Duration = "monthly"
	TwoMonth   Duration = "two_month"
	ThreeMonth Duration = "three_month"
	SixMonth   Duration = "six_month"
	Yearly     Duration = "yearly"
	Lifetime   Duration = "lifetime"
)

// toMilliseconds takes a time and returns Unix epoch in milliseconds.
func toMilliseconds(t time.Time) int64 {
	return t.UTC().UnixNano() / 1e6
}

// fromMilliseconds takes a Unix epoch in milliseconds value and returns a time.Time.
func fromMilliseconds(t int64) time.Time {
	return time.Unix(0, t*1e6)
}

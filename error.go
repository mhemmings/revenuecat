package revenuecat

import "fmt"

// Error represents an error returned by RevenueCat
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err Error) Error() string {
	return fmt.Sprintf("%d: %s", err.Code, err.Message)
}

//go:generate paramgen -output=paramagen.go Config

package source

import "time"

type Config struct {
	URL              string        `json:"url"`
	Name             string        `json:"name"`
	SubscriptionType string        `json:"subscriptionType" default:"exclusive" validate:"inclusion=exclusive|shared|failover"`
	Topic            string        `json:"topic"`
	AckGroupMaxTime  time.Duration `json:"ackGroundMaxTime"`
	AckGroupMaxSize  uint32        `json:"ackGroupMaxSize"`
}

//go:generate paramgen -output=paramagen.go Config

package source

import "time"

type Config struct {
	URL              string        `json:"url" default:"pulsar://localhost:6650"`
	Name             string        `json:"name" default:"conduit-connector"`
	SubscriptionType string        `json:"subscriptionType" default:"exclusive" validate:"inclusion=exclusive|shared|failover"`
	Topic            string        `json:"topic" validate:"required"`
	AckGroupMaxTime  time.Duration `json:"ackGroundMaxTime"`
	AckGroupMaxSize  uint32        `json:"ackGroupMaxSize"`
}

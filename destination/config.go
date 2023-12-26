//go:generate paramgen -output=paramagen.go Config

package destination

import (
	"time"
)

type Config struct {
	Topic                   string        `json:"topic"`
	Name                    string        `json:"name"`
	Url                     string        `json:"url" `
	SendTimeout             time.Duration `json:"sendTimeout"`
	DisableBlockIfQueueFull bool          `json:"disableBlockIfQueueFull" validate:"inclusion=true|false"`
	MaxPendingMessages      int           `json:"maxPendingMessages"`
	//HashingScheme           string        `json:"hashingScheme" default:"java-string-hash" validate:"inclusion=java-string-hash|murmur3-32"`
	//CompressionType string `json:"compressionType" default:"lz4" validate:"inclusion=lz4|zlib|zstd|6none"`
}

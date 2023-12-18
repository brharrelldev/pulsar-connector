package source

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
)

type Consumer interface {
	Consume(ctx context.Context) (pulsar.Message, error)
	Close(ctx context.Context) error
}

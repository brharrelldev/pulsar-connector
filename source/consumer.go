package source

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
)

////go:generate mockgen -destination consumer_mock.go -package source -mock_names=Consumer=MockConsumer . Consumer

type Consumer interface {
	Consume(ctx context.Context) (Record, error)
	Close(ctx context.Context) error
}

type Record pulsar.Message

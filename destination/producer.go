package destination

//go:generate mockgen -destination mock_producer.go -package destination -mock_names=Producer=MockProducer . Producer

import (
	"context"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Producer interface {
	Produce(ctx context.Context, records []sdk.Record) (int, error)
	Close(ctx context.Context) error
}

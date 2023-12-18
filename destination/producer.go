package destination

import (
	"context"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Producer interface {
	Produce(ctx context.Context, records []sdk.Record) (int, error)
	Close(ctx context.Context) error
}

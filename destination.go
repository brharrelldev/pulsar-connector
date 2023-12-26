package pulsar

import (
	"context"
	"github.com/brharrelldev/pulsar/destination"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Destination struct {
	sdk.UnimplementedDestination
	producer destination.Producer
	config   destination.Config
}

func NewDestination() sdk.Destination {
	return sdk.DestinationWithMiddleware(&Destination{}, sdk.DefaultDestinationMiddleware()...)
}

func (d *Destination) Parameters() map[string]sdk.Parameter {
	return destination.Config{}.Parameters()

}

func (d *Destination) Configure(ctx context.Context, cfg map[string]string) error {

	var dc destination.Config
	if err := sdk.Util.ParseConfig(cfg, &dc); err != nil {
		return err
	}

	d.config = dc

	return nil
}

func (d *Destination) Open(ctx context.Context) error {

	producer, err := destination.NewPulsarProducer(ctx, d.config)
	if err != nil {
		return err
	}

	d.producer = producer

	return nil
}

func (d *Destination) Write(ctx context.Context, records []sdk.Record) (int, error) {
	return d.producer.Produce(ctx, records)
}

func (d *Destination) Teardown(ctx context.Context) error {

	if d.producer != nil {
		if err := d.producer.Close(ctx); err != nil {
			return err
		}
	}

	return nil
}

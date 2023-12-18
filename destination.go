package main

import (
	"context"
	"github.com/brharrelldev/pulsar_connector/destination"
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

func (d *Destination) Configure(ctx context.Context, cfg map[string]string) error {

	var dc destination.Config
	if err := sdk.Util.ParseConfig(cfg, dc); err != nil {
		return err
	}

	d.config = dc

	return nil
}

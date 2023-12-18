package main

import (
	"context"
	"github.com/brharrelldev/pulsar_connector/source"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Source struct {
	sdk.UnimplementedSource
	source source.Consumer
	config source.Config
}

func NewSource() sdk.Source {
	return sdk.SourceWithMiddleware(&Source{}, sdk.DefaultSourceMiddleware()...)

}

func (s *Source) Configure(ctx context.Context, cfg map[string]string) error {

	var sc source.Config
	if err := sdk.Util.ParseConfig(cfg, sc); err != nil {
		return err
	}

	s.config = sc

	return nil
}

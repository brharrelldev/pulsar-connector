package pulsar

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brharrelldev/pulsar/source"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Source struct {
	sdk.UnimplementedSource
	source source.Consumer
	config source.Config
}

type Position struct {
	Offset int
	Cursor int
}

func NewSource() sdk.Source {
	return sdk.SourceWithMiddleware(&Source{}, sdk.DefaultSourceMiddleware()...)

}

func (s *Source) Parameters() map[string]sdk.Parameter {
	return source.Config{}.Parameters()

}

func (s *Source) Configure(ctx context.Context, cfg map[string]string) error {

	var sc source.Config
	if err := sdk.Util.ParseConfig(cfg, &sc); err != nil {
		return err
	}

	sdk.Logger(ctx).Info().Msgf("Config: %+v", sc)

	s.config = sc

	return nil
}

func (s *Source) Open(ctx context.Context, sdkPos sdk.Position) error {

	sdk.Logger(ctx).Info().Msgf("inspecing config %v", s.config)
	consumer, err := source.NewPulsarConsumer(ctx, s.config)
	if err != nil {
		return fmt.Errorf("error creating pulsar consumer: %v", err)
	}

	s.source = consumer

	return nil

}

func (s *Source) Read(ctx context.Context) (sdk.Record, error) {

	msg, err := s.source.Consume(ctx)
	if err != nil {
		return sdk.Record{}, fmt.Errorf("error consuming message: %v", err)
	}

	metadata := sdk.Metadata{"topic": msg.Topic()}
	metadata.SetCreatedAt(*msg.BrokerPublishTime())

	pos := Position{}

	posBytes, err := json.Marshal(pos)
	if err != nil {
		return sdk.Record{}, fmt.Errorf("error marshalling position: %v", err)
	}

	return sdk.Util.Source.NewRecordCreate(posBytes, metadata, sdk.RawData(msg.Key()), sdk.RawData(msg.Payload())), nil

}

func (s *Source) Teardown(ctx context.Context) error {

	if s.source != nil {
		if err := s.source.Close(ctx); err != nil {
			return err
		}
	}

	return nil

}

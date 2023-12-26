package destination

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	sdk "github.com/conduitio/conduit-connector-sdk"
	"log"
)

type PulsarProducer struct {
	producer pulsar.Producer
	encoder  dataEncoder
}

func NewPulsarProducer(ctx context.Context, cfg Config) (*PulsarProducer, error) {
	pc, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: cfg.Url,
	})

	log.Println("Pulsar Client: ", cfg.Url)

	if err != nil {
		return nil, fmt.Errorf("error creating pulsar client: %v", err)
	}

	producer, err := pc.CreateProducer(pulsar.ProducerOptions{
		Topic: cfg.Topic,
		Name:  cfg.Name,
	})

	if err != nil {
		return nil, fmt.Errorf("error creating pulsar producer: %v", err)
	}

	de := dataEncoder{}

	return &PulsarProducer{
		producer: producer,
		encoder:  de,
	}, nil

}

func (p *PulsarProducer) Produce(ctx context.Context, records []sdk.Record) (int, error) {

	count := 0
	for _, rec := range records {
		encodeKey, err := p.encoder.Encode(rec.Key)
		if err != nil {
			return 0, fmt.Errorf("error encoding key: %v", err)
		}

		if _, err := p.producer.Send(ctx, &pulsar.ProducerMessage{
			Key:   string(encodeKey),
			Value: rec.Bytes(),
		}); err != nil {
			return 0, fmt.Errorf("error sending message: %v", err)
		}

		count++
	}

	return count, nil

}

func (p *PulsarProducer) Close(ctx context.Context) error {

	p.producer.Close()

	return nil

}

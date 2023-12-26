package source

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
)

type PulsarConsumer struct {
	consumer pulsar.Consumer
}

var _ Consumer = (*PulsarConsumer)(nil)

func NewPulsarConsumer(ctx context.Context, cfg Config) (*PulsarConsumer, error) {

	log.Println("Pulsar Client: ", cfg.URL)
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: cfg.URL,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating pulsar client: %v url: %s", err, cfg.URL)
	}

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            cfg.Topic,
		SubscriptionName: cfg.Name,
		Type:             ConvertSubscriptionTypes(cfg.SubscriptionType),
		AckGroupingOptions: &pulsar.AckGroupingOptions{
			MaxSize: cfg.AckGroupMaxSize,
			MaxTime: cfg.AckGroupMaxTime,
		},
	})
	if err != nil {
		return nil, err
	}

	return &PulsarConsumer{
		consumer: consumer,
	}, nil

}

func (c *PulsarConsumer) Consume(ctx context.Context) (Record, error) {

	msg, err := c.consumer.Receive(context.Background())
	if err != nil {
		return nil, err
	}

	return msg, nil

}

func (c *PulsarConsumer) Close(ctx context.Context) error {

	err := c.consumer.Unsubscribe()
	if err != nil {
		return err
	}
	c.consumer.Close()

	return nil

}

func ConvertSubscriptionTypes(name string) pulsar.SubscriptionType {
	switch name {
	case "exclusive":
		return pulsar.Exclusive
	case "shared":
		return pulsar.Shared
	case "failover":
		return pulsar.Failover
	case "key_shared":
		return pulsar.KeyShared
	default:
		return pulsar.Exclusive
	}

}

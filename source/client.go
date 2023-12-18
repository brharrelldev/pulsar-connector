package source

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
)

type PulsarConsumer struct {
	consumer pulsar.Consumer
}

func NewPulsarConsumer(ctx context.Context, cfg Config) (*PulsarConsumer, error) {

	client, err := pulsar.NewClient(pulsar.ClientOptions{})
	if err != nil {
		return nil, err
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

func (c *PulsarConsumer) Consumer() (pulsar.Message, error) {

	msg, err := c.consumer.Receive(context.Background())
	if err != nil {
		return nil, err
	}

	return msg, nil

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

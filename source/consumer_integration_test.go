package source

import (
	"context"
	"fmt"
	sdk "github.com/conduitio/conduit-connector-sdk"
	"testing"
)

func Test_NewConsumer(t *testing.T) {

	configMap := map[string]string{
		"URL":   "pulsar://localhost:6650",
		"topic": "test",
		"name":  "test",
	}

	var c Config

	if err := sdk.Util.ParseConfig(configMap, &c); err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", c)

	consumer, err := NewPulsarConsumer(context.Background(), c)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", consumer)

}

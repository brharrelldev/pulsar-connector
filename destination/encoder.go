package destination

import sdk "github.com/conduitio/conduit-connector-sdk"

type dataEncoder struct {
}

type Encoder interface {
	Encode(sdk.Data) ([]byte, error)
}

func (d *dataEncoder) Encode(data sdk.Data) ([]byte, error) {
	return data.Bytes(), nil
}

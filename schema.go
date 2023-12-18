package main

import sdk "github.com/conduitio/conduit-connector-sdk"

var Version string

func NewSchema() sdk.Specification {

	return sdk.Specification{
		Name:        "pulsar-connector",
		Summary:     "alpha pulsar connector",
		Description: "alpha pulsar connector",
		Version:     Version,
		Author:      "",
	}

}

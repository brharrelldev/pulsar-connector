package main

import (
	pulsar_connect "github.com/brharrelldev/pulsar"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func main() {

	sdk.Serve(pulsar_connect.Connector)
}

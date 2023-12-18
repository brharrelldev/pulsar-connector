package main

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

var Connector = sdk.Connector{
	NewDestination:   NewDestination,
	NewSource:        NewSource,
	NewSpecification: NewSchema,
}

func main() {

	sdk.Serve(Connector)
}

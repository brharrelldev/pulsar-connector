// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-connector-sdk/tree/main/cmd/paramgen

package destination

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func (Config) Parameters() map[string]sdk.Parameter {
	return map[string]sdk.Parameter{
		"compressionType": {
			Default:     "lz4",
			Description: "",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationInclusion{List: []string{"lz4", "zlib", "zstd", "6none"}},
			},
		},
		"disableBlockIfQueueFull": {
			Default:     "",
			Description: "",
			Type:        sdk.ParameterTypeBool,
			Validations: []sdk.Validation{
				sdk.ValidationInclusion{List: []string{"true", "false"}},
			},
		},
		"hashingScheme": {
			Default:     "java-string-hash",
			Description: "",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationInclusion{List: []string{"java-string-hash", "murmur3-32"}},
			},
		},
		"maxPendingMessages": {
			Default:     "",
			Description: "",
			Type:        sdk.ParameterTypeInt,
			Validations: []sdk.Validation{},
		},
		"name": {
			Default:     "",
			Description: "",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
		"sendTimeout": {
			Default:     "",
			Description: "",
			Type:        sdk.ParameterTypeDuration,
			Validations: []sdk.Validation{},
		},
		"topic": {
			Default:     "",
			Description: "",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{},
		},
	}
}
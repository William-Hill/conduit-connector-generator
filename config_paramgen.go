// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-commons/tree/main/paramgen

package generator

import (
	"github.com/conduitio/conduit-commons/config"
)

const (
	ConfigBurstGenerateTime            = "burst.generateTime"
	ConfigBurstSleepTime               = "burst.sleepTime"
	ConfigCollectionsFormatOptions     = "collections.*.format.options.*"
	ConfigCollectionsFormatOptionsPath = "collections.*.format.options.path"
	ConfigCollectionsFormatType        = "collections.*.format.type"
	ConfigCollectionsOperations        = "collections.*.operations"
	ConfigFormatOptions                = "format.options.*"
	ConfigFormatOptionsPath            = "format.options.path"
	ConfigFormatType                   = "format.type"
	ConfigOperations                   = "operations"
	ConfigRate                         = "rate"
	ConfigReadTime                     = "readTime"
	ConfigRecordCount                  = "recordCount"
)

func (Config) Parameters() map[string]config.Parameter {
	return map[string]config.Parameter{
		ConfigBurstGenerateTime: {
			Default:     "1s",
			Description: "The amount of time the generator is generating records in a burst. Has an\neffect only if `burst.sleepTime` is set.",
			Type:        config.ParameterTypeDuration,
			Validations: []config.Validation{},
		},
		ConfigBurstSleepTime: {
			Default:     "",
			Description: "The time the generator \"sleeps\" between bursts.",
			Type:        config.ParameterTypeDuration,
			Validations: []config.Validation{},
		},
		ConfigCollectionsFormatOptions: {
			Default:     "",
			Description: "The options for the `raw` and `structured` format types. It accepts pairs\nof field names and field types, where the type can be one of: `int`, `string`, `time`, `bool`, `duration`,\n`name`, `email`, `employeeid`, `ssn`, `creditcard`, `ordernumber`.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		ConfigCollectionsFormatOptionsPath: {
			Default:     "",
			Description: "Path to the input file (only applicable if the format type is `file`).",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		ConfigCollectionsFormatType: {
			Default:     "",
			Description: "The format of the generated payload data (raw, structured, file, fhir, hl7).",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationInclusion{List: []string{"raw", "structured", "file", "fhir", "hl7"}},
			},
		},
		ConfigCollectionsOperations: {
			Default:     "create",
			Description: "Comma separated list of record operations to generate. Allowed values are\n\"create\", \"update\", \"delete\", \"snapshot\".",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		ConfigFormatOptions: {
			Default:     "",
			Description: "The options for the `raw` and `structured` format types. It accepts pairs\nof field names and field types, where the type can be one of: `int`, `string`, `time`, `bool`, `duration`,\n`name`, `email`, `employeeid`, `ssn`, `creditcard`, `ordernumber`.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		ConfigFormatOptionsPath: {
			Default:     "",
			Description: "Path to the input file (only applicable if the format type is `file`).",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		ConfigFormatType: {
			Default:     "",
			Description: "The format of the generated payload data (raw, structured, file, fhir, hl7).",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationInclusion{List: []string{"raw", "structured", "file", "fhir", "hl7"}},
			},
		},
		ConfigOperations: {
			Default:     "create",
			Description: "Comma separated list of record operations to generate. Allowed values are\n\"create\", \"update\", \"delete\", \"snapshot\".",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		ConfigRate: {
			Default:     "",
			Description: "The maximum rate in records per second, at which records are generated (0\nmeans no rate limit).",
			Type:        config.ParameterTypeFloat,
			Validations: []config.Validation{},
		},
		ConfigReadTime: {
			Default:     "",
			Description: "The time it takes to 'read' a record.\nDeprecated: use `rate` instead.",
			Type:        config.ParameterTypeDuration,
			Validations: []config.Validation{},
		},
		ConfigRecordCount: {
			Default:     "",
			Description: "Number of records to be generated (0 means infinite).",
			Type:        config.ParameterTypeInt,
			Validations: []config.Validation{
				config.ValidationGreaterThan{V: -1},
			},
		},
	}
}

package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FraudDetectorResources() []*Resource {
	arnColumn := codegen.ColumnDefinition{
			Name:     "arn",
			Type:     schema.TypeString,
			Resolver: `schema.PathResolver("Arn")`,
			Options:  schema.ColumnCreationOptions{PrimaryKey: true},
	}
	skipARN := "Arn"
	resources := []*Resource{
		{
			SubService: "batch_import",
			Struct:     new(types.BatchImport),
		},
		{
			SubService: "batch_prediction",
			Struct:     new(types.BatchPrediction),
		},
		{
			SubService: "detector",
			Struct:     new(types.Detector),
		},
		{
			SubService: "entity_type",
			Struct:     new(types.EntityType),
		},
		{
			SubService: "event_type",
			Struct:     new(types.EventType),
		},
		{
			SubService: "external_model",
			Struct:     new(types.ExternalModel),
		},
		{
			SubService: "label",
			Struct:     new(types.Label),
		},
		{
			SubService: "model",
			Struct:     new(types.Model),
		},
		{
			SubService: "model_version",
			Struct:     new(types.ModelVersion),
		},
		{
			SubService: "model_version_detail",
			Struct:     new(types.ModelVersionDetail),
		},
		{
			SubService: "outcome",
			Struct:     new(types.Outcome),
		},
		{
			SubService: "rule_detail",
			Struct:     new(types.RuleDetail),
		},
		{
			SubService: "variable",
			Struct:     new(types.Variable),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "frauddetector"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("frauddetector")`
		r.ExtraColumns = append(r.ExtraColumns, arnColumn)
		r.SkipFields = append(r.SkipFields, skipARN)
	}
	return resources
}

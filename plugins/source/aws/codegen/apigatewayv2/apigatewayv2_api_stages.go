// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

func Apigatewayv2ApiStages() *schema.Table {
	return &schema.Table{
		Name:      "aws_apigatewayv2_api_stages",
		Resolver:  fetchApigatewayv2ApiStages,
		Multiplex: client.ServiceAccountRegionMultiplexer("apigatewayv2"),
		Columns: []schema.Column{
			{
				Name:     "stage_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StageName"),
			},
			{
				Name:     "access_log_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AccessLogSettings"),
			},
			{
				Name:     "api_gateway_managed",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ApiGatewayManaged"),
			},
			{
				Name:     "auto_deploy",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoDeploy"),
			},
			{
				Name:     "client_certificate_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientCertificateId"),
			},
			{
				Name:     "created_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "default_route_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultRouteSettings"),
			},
			{
				Name:     "deployment_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeploymentId"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "last_deployment_status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastDeploymentStatusMessage"),
			},
			{
				Name:     "last_updated_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedDate"),
			},
			{
				Name:     "route_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RouteSettings"),
			},
			{
				Name:     "stage_variables",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StageVariables"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}

func fetchApigatewayv2ApiStages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2

	r1 := parent.Item.(types.Api)

	input := apigatewayv2.GetStagesInput{
		ApiId: r1.ApiId,
	}

	for {
		response, err := svc.GetStages(ctx, &input)
		if err != nil {

			return diag.WrapError(err)
		}

		res <- response.Items

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

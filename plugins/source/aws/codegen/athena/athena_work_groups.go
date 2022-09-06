// Code generated by codegen using template resource_list_and_detail.go.tpl; DO NOT EDIT.

package athena

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	resolvers "github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/resolvers/athena"
)

func AthenaWorkGroups() *schema.Table {
	return &schema.Table{
		Name:      "aws_athena_work_groups",
		Resolver:  fetchAthenaWorkGroups,
		Multiplex: client.ServiceAccountRegionMultiplexer("athena"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
			},
			{
				Name:        "region",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
				Description: `The AWS Region of the resource.`,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Configuration"),
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolvers.ResolveWorkGroupArn,
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveAthenaWorkGroupsTags,
				Description: `Tags associated with the Athena work group.`,
			},
		},

		Relations: []*schema.Table{
			AthenaWorkGroupPreparedStatements(),
			AthenaWorkGroupQueryExecutions(),
			AthenaWorkGroupNamedQueries(),
		},
	}
}

func fetchAthenaWorkGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return errors.WithStack(client.ListAndDetailResolver(ctx, meta, res, listWorkGroups, listWorkGroupsDetail))
}

func listWorkGroups(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena

	input := athena.ListWorkGroupsInput{}

	for {
		response, err := svc.ListWorkGroups(ctx, &input)
		if err != nil {
			return errors.WithStack(err)
		}
		for _, item := range response.WorkGroups {
			detailChan <- item
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func listWorkGroupsDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	cl := meta.(*client.Client)
	item := listInfo.(types.WorkGroupSummary)
	svc := cl.Services().Athena
	response, err := svc.GetWorkGroup(ctx, &athena.GetWorkGroupInput{
		WorkGroup: item.Name,
	})
	if err != nil {

		if cl.IsNotFoundError(err) {
			return
		}
		errorChan <- errors.WithStack(err)
		return
	}
	resultsChan <- *response.WorkGroup
}

func resolveAthenaWorkGroupsTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	item := resource.Item.(types.WorkGroup)
	params := athena.ListTagsForResourceInput{
		ResourceARN: aws.String(resolvers.CreateWorkGroupArn(cl, *item.Name)),
	}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return errors.WithStack(err)
		}
		client.TagsIntoMap(result.Tags, tags)
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return errors.WithStack(resource.Set(c.Name, tags))
}

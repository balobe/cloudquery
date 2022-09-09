package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:        "aws_xray_groups",
		Description: "Details for a group without metadata",
		Resolver:    fetchXrayGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("xray"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveXrayGroupTags,
			},
			{
				Name:        "filter_expression",
				Description: "The filter expression defining the parameters to include traces",
				Type:        schema.TypeString,
			},
			{
				Name:            "arn",
				Description:     "The ARN of the group generated based on the GroupName",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("GroupARN"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "group_name",
				Description: "The unique case-sensitive name of the group",
				Type:        schema.TypeString,
			},
			{
				Name:     "insights_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InsightsConfiguration"),
			},
			{
				Name:     "insights_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InsightsConfiguration"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchXrayGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := xray.NewGetGroupsPaginator(meta.(*client.Client).Services().Xray, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.Groups
	}
	return nil
}
func resolveXrayGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	group := resource.Item.(types.GroupSummary)
	cl := meta.(*client.Client)
	svc := cl.Services().Xray
	params := xray.ListTagsForResourceInput{ResourceARN: group.GroupARN}

	output, err := svc.ListTagsForResource(ctx, &params)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	tags := map[string]string{}
	client.TagsIntoMap(output.Tags, tags)

	return resource.Set(c.Name, tags)
}

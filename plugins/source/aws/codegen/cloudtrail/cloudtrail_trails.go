// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package cloudtrail

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
)

func CloudtrailTrails() *schema.Table {
	return &schema.Table{
		Name:      "aws_cloudtrail_trails",
		Resolver:  fetchCloudtrailTrails,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
			},
			{
				Name:     "cloud_watch_logs_log_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CloudWatchLogsLogGroupArn"),
			},
			{
				Name:     "cloud_watch_logs_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CloudWatchLogsRoleArn"),
			},
			{
				Name:     "has_custom_event_selectors",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HasCustomEventSelectors"),
			},
			{
				Name:     "has_insight_selectors",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HasInsightSelectors"),
			},
			{
				Name:     "home_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HomeRegion"),
			},
			{
				Name:     "include_global_service_events",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IncludeGlobalServiceEvents"),
			},
			{
				Name:     "is_multi_region_trail",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsMultiRegionTrail"),
			},
			{
				Name:     "is_organization_trail",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsOrganizationTrail"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "log_file_validation_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("LogFileValidationEnabled"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "s_3_bucket_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("S3BucketName"),
			},
			{
				Name:     "s_3_key_prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("S3KeyPrefix"),
			},
			{
				Name:     "sns_topic_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnsTopicARN"),
			},
			{
				Name:     "sns_topic_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnsTopicName"),
			},
			{
				Name:     "trail_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TrailARN"),
			},
		},

		Relations: []*schema.Table{
			CloudtrailTrailEventSelectors(),
		},
	}
}

func fetchCloudtrailTrails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudtrail

	input := cloudtrail.DescribeTrailsInput{}

	{
		response, err := svc.DescribeTrails(ctx, &input)
		if err != nil {

			return errors.WithStack(err)
		}

		res <- response.TrailList

	}
	return nil
}

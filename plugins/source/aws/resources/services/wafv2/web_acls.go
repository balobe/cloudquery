package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type WebACLWrapper struct {
	*types.WebACL
	LoggingConfiguration *types.LoggingConfiguration
}

func Wafv2WebAcls() *schema.Table {
	return &schema.Table{
		Name:        "aws_wafv2_web_acls",
		Description: "A Web ACL defines a collection of rules to use to inspect and control web requests",
		Resolver:    fetchWafv2WebAcls,
		Multiplex:   client.ServiceAccountRegionScopeMultiplexer("waf-regional"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:        "scope",
				Description: "Specifies whether this is for an Amazon CloudFront distribution or for a regional application.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveWAFScope,
			},
			{
				Name:     "resources_for_web_acl",
				Type:     schema.TypeStringArray,
				Resolver: resolveWafv2webACLResourcesForWebACL,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveWafv2webACLTags,
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.  ",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "default_action",
				Description: "The action to perform if none of the Rules contained in the WebACL match.  ",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("DefaultAction"),
			},
			{
				Name:        "id",
				Description: "A unique identifier for the WebACL",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "name",
				Description: "The name of the Web ACL",
				Type:        schema.TypeString,
			},
			{
				Name:     "visibility_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VisibilityConfig"),
			},
			{
				Name:     "visibility_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VisibilityConfig"),
			},
			{
				Name:     "visibility_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VisibilityConfig"),
			},
			{
				Name:        "capacity",
				Description: "The web ACL capacity units (WCUs) currently being used by this web ACL",
				Type:        schema.TypeInt,
			},
			{
				Name:        "custom_response_bodies",
				Description: "A map of custom response keys and content bodies",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "description",
				Description: "A description of the Web ACL that helps with identification.",
				Type:        schema.TypeString,
			},
			{
				Name:        "label_namespace",
				Description: "The label namespace prefix for this web ACL",
				Type:        schema.TypeString,
			},
			{
				Name:        "managed_by_firewall_manager",
				Description: "Indicates whether this web ACL is managed by AWS Firewall Manager",
				Type:        schema.TypeBool,
			},
			{
				Name:     "logging_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LoggingConfiguration"),
			},
			{
				Name:        "rules",
				Description: "A single rule, which you can use in a WebACL or RuleGroup to identify web requests that you want to allow, block, or count",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Rules"),
			},
			{
				Name: "post_process_firewall_manager_rule_groups",
				Type: schema.TypeJSON,
			},
			{
				Name: "pre_process_firewall_manager_rule_groups",
				Type: schema.TypeJSON,
			},
			{
				Name: "logging_configuration",
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchWafv2WebAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	service := c.Services().WafV2

	config := wafv2.ListWebACLsInput{
		Scope: c.WAFScope,
		Limit: aws.Int32(100),
	}
	for {
		output, err := service.ListWebACLs(ctx, &config)
		if err != nil {
			return err
		}
		for _, webAcl := range output.WebACLs {
			webAclConfig := wafv2.GetWebACLInput{Id: webAcl.Id, Name: webAcl.Name, Scope: c.WAFScope}
			webAclOutput, err := service.GetWebACL(ctx, &webAclConfig, func(options *wafv2.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}

			cfg := wafv2.GetLoggingConfigurationInput{
				ResourceArn: webAclOutput.WebACL.ARN,
			}

			loggingConfigurationOutput, err := service.GetLoggingConfiguration(ctx, &cfg, func(options *wafv2.Options) {
				options.Region = c.Region
			})
			if err != nil {
				if client.IsAWSError(err, "WAFNonexistentItemException") {
					c.Logger().Debug().Err(err).Msg("Logging configuration not found for")
				} else {
					c.Logger().Error().Err(err).Msg("GetLoggingConfiguration failed with error")
				}
			}

			var webAclLoggingConfiguration *types.LoggingConfiguration
			if loggingConfigurationOutput != nil {
				webAclLoggingConfiguration = loggingConfigurationOutput.LoggingConfiguration
			}

			res <- &WebACLWrapper{
				webAclOutput.WebACL,
				webAclLoggingConfiguration,
			}
		}

		if aws.ToString(output.NextMarker) == "" {
			break
		}
		config.NextMarker = output.NextMarker
	}
	return nil
}
func resolveWafv2webACLResourcesForWebACL(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL := resource.Item.(*WebACLWrapper)

	cl := meta.(*client.Client)
	service := cl.Services().WafV2

	resourceArns := []string{}
	if cl.WAFScope == types.ScopeCloudfront {
		cloudfrontService := cl.Services().Cloudfront
		params := &cloudfront.ListDistributionsByWebACLIdInput{
			WebACLId: webACL.Id,
			MaxItems: aws.Int32(100),
		}
		for {
			output, err := cloudfrontService.ListDistributionsByWebACLId(ctx, params, func(options *cloudfront.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			for _, item := range output.DistributionList.Items {
				resourceArns = append(resourceArns, *item.ARN)
			}
			if aws.ToString(output.DistributionList.NextMarker) == "" {
				break
			}
			params.Marker = output.DistributionList.NextMarker
		}
	} else {
		output, err := service.ListResourcesForWebACL(ctx, &wafv2.ListResourcesForWebACLInput{WebACLArn: webACL.ARN})
		if err != nil {
			return err
		}
		resourceArns = output.ResourceArns
	}
	return resource.Set(c.Name, resourceArns)
}
func resolveWafv2webACLTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	webACL := resource.Item.(*WebACLWrapper)

	cl := meta.(*client.Client)
	service := cl.Services().WafV2

	// Resolve tags
	outputTags := make(map[string]*string)
	tagsConfig := wafv2.ListTagsForResourceInput{ResourceARN: webACL.ARN}
	for {
		tags, err := service.ListTagsForResource(ctx, &tagsConfig)
		if err != nil {
			return err
		}
		for _, t := range tags.TagInfoForResource.TagList {
			outputTags[*t.Key] = t.Value
		}
		if aws.ToString(tags.NextMarker) == "" {
			break
		}
		tagsConfig.NextMarker = tags.NextMarker
	}
	return resource.Set(c.Name, outputTags)
}

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/elasticbeanstalk/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchElasticbeanstalkEnvironments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config elasticbeanstalk.DescribeEnvironmentsInput
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk
	for {
		response, err := svc.DescribeEnvironments(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Environments
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveElasticbeanstalkEnvironmentTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.EnvironmentDescription)
	if p.Resources == nil || p.Resources.LoadBalancer == nil {
		return nil
	}
	listeners := make(map[int32]*string, len(p.Resources.LoadBalancer.Listeners))
	for _, l := range p.Resources.LoadBalancer.Listeners {
		listeners[l.Port] = l.Protocol
	}
	return resource.Set(c.Name, listeners)
}
func resolveElasticbeanstalkEnvironmentListeners(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.EnvironmentDescription)
	cl := meta.(*client.Client)
	svc := cl.Services().ElasticBeanstalk
	tagsOutput, err := svc.ListTagsForResource(ctx, &elasticbeanstalk.ListTagsForResourceInput{
		ResourceArn: p.EnvironmentArn,
	}, func(o *elasticbeanstalk.Options) {})
	if err != nil {
		// It takes a few minutes for an environment to be terminated
		// This ensures we don't error while trying to fetch related resources for a terminated environment
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if len(tagsOutput.ResourceTags) == 0 {
		return nil
	}
	tags := make(map[string]*string)
	for _, s := range tagsOutput.ResourceTags {
		tags[*s.Key] = s.Value
	}
	return resource.Set(c.Name, tags)
}

func fetchElasticbeanstalkConfigurationOptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(types.EnvironmentDescription)
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk
	configOptionsIn := elasticbeanstalk.DescribeConfigurationOptionsInput{
		ApplicationName: p.ApplicationName,
		EnvironmentName: p.EnvironmentName,
	}
	output, err := svc.DescribeConfigurationOptions(ctx, &configOptionsIn)
	if err != nil {
		// It takes a few minutes for an environment to be terminated
		// This ensures we don't error while trying to fetch related resources for a terminated environment
		if client.IsInvalidParameterValueError(err) {
			meta.Logger().Debug().Interface("environment", p.EnvironmentName).Interface("application", p.ApplicationName).Msg("Failed extracting configuration options for environment. It might be terminated")
			return nil
		}
		return err
	}

	for _, option := range output.Options {
		res <- models.ConfigurationOptionDescriptionWrapper{
			ConfigurationOptionDescription: option, ApplicationArn: c.ARN("elasticbeanstalk", "application", *p.ApplicationName),
		}
	}

	return nil
}

func fetchElasticbeanstalkConfigurationSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(types.EnvironmentDescription)
	c := meta.(*client.Client)
	svc := c.Services().ElasticBeanstalk

	configOptionsIn := elasticbeanstalk.DescribeConfigurationSettingsInput{
		ApplicationName: p.ApplicationName,
		EnvironmentName: p.EnvironmentName,
	}
	output, err := svc.DescribeConfigurationSettings(ctx, &configOptionsIn)
	if err != nil {
		// It takes a few minutes for an environment to be terminated
		// This ensures we don't error while trying to fetch related resources for a terminated environment
		if client.IsInvalidParameterValueError(err) {
			meta.Logger().Debug().Interface("environment", p.EnvironmentName).Interface("application", p.ApplicationName).Msg("Failed extracting configuration settings for environment. It might be terminated")
			return nil
		}
		return err
	}

	for _, option := range output.ConfigurationSettings {
		res <- models.ConfigurationSettingsDescriptionWrapper{
			ConfigurationSettingsDescription: option, ApplicationArn: c.ARN("elasticbeanstalk", "application", *p.ApplicationName),
		}
	}

	return nil
}

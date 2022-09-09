package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Templates() *schema.Table {
	return &schema.Table{
		Name:        "aws_ses_templates",
		Description: "Amazon Simple Email Service (SES) is a cost-effective, flexible, and scalable email service that enables developers to send mail from within any application.",
		Resolver:    fetchSesTemplates,
		Multiplex:   client.ServiceAccountRegionMultiplexer("email"),
		Columns: []schema.Column{
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) for the resource.",
				Type:            schema.TypeString,
				Resolver:        ResolveSesTemplateArn,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
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
				Name:        "name",
				Description: "The name of the template.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TemplateName"),
			},
			{
				Name:     "email_template_content",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EmailTemplateContent"),
			},
			{
				Name:     "email_template_content",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EmailTemplateContent"),
			},
			{
				Name:     "email_template_content",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EmailTemplateContent"),
			},
			{
				Name:        "created_timestamp",
				Description: "The time and date the template was created.",
				Type:        schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSesTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().SES

	listInput := new(sesv2.ListEmailTemplatesInput)
	for {
		output, err := svc.ListEmailTemplates(ctx, listInput, func(o *sesv2.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}

		for _, templateMeta := range output.TemplatesMetadata {
			getInput := &sesv2.GetEmailTemplateInput{TemplateName: templateMeta.TemplateName}
			getOutput, err := svc.GetEmailTemplate(ctx, getInput, func(o *sesv2.Options) {
				o.Region = c.Region
			})
			if err != nil {
				return err
			}

			res <- &Template{
				TemplateName:         getOutput.TemplateName,
				CreatedTimestamp:     templateMeta.CreatedTimestamp,
				EmailTemplateContent: getOutput.TemplateContent,
			}
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}
		listInput.NextToken = output.NextToken
	}

	return nil
}
func ResolveSesTemplateArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return client.ResolveARN(client.SESService, func(resource *schema.Resource) ([]string, error) {
		return []string{"template", *resource.Item.(*Template).TemplateName}, nil
	})(ctx, meta, resource, c)
}

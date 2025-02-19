// Code generated by codegen; DO NOT EDIT.

package efs

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Filesystems() *schema.Table {
	return &schema.Table{
		Name:        "aws_efs_filesystems",
		Description: "https://docs.aws.amazon.com/efs/latest/ug/API_FileSystemDescription.html",
		Resolver:    fetchEfsFilesystems,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticfilesystem"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FileSystemArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "backup_policy_status",
				Type:     schema.TypeString,
				Resolver: ResolveEfsFilesystemBackupPolicyStatus,
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "creation_token",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationToken"),
			},
			{
				Name:     "file_system_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FileSystemId"),
			},
			{
				Name:     "life_cycle_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LifeCycleState"),
			},
			{
				Name:     "number_of_mount_targets",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumberOfMountTargets"),
			},
			{
				Name:     "owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerId"),
			},
			{
				Name:     "performance_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PerformanceMode"),
			},
			{
				Name:     "size_in_bytes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SizeInBytes"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "availability_zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZoneId"),
			},
			{
				Name:     "availability_zone_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZoneName"),
			},
			{
				Name:     "encrypted",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Encrypted"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "provisioned_throughput_in_mibps",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("ProvisionedThroughputInMibps"),
			},
			{
				Name:     "throughput_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ThroughputMode"),
			},
		},
	}
}

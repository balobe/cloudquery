// Code generated by codegen; DO NOT EDIT.

package snapshots

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_snapshots",
		Resolver: fetchSnapshotsSnapshots,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceID"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "regions",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Regions"),
			},
			{
				Name:     "min_disk_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinDiskSize"),
			},
			{
				Name:     "size_gigabytes",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("SizeGigaBytes"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}

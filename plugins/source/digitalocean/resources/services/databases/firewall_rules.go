// Code generated by codegen; DO NOT EDIT.

package databases

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func FirewallRules() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_database_firewall_rules",
		Resolver: fetchDatabasesFirewallRules,
		Columns: []schema.Column{
			{
				Name:     "uuid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UUID"),
			},
			{
				Name:     "cluster_uuid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterUUID"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "value",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Value"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
		},
	}
}

// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package fsx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceONTAPFileSystem,
			TypeName: "aws_fsx_ontap_file_system",
			Name:     "ONTAP File System",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceONTAPStorageVirtualMachine,
			TypeName: "aws_fsx_ontap_storage_virtual_machine",
			Name:     "ONTAP Storage Virtual Machine",
		},
		{
			Factory:  dataSourceONTAPStorageVirtualMachines,
			TypeName: "aws_fsx_ontap_storage_virtual_machines",
			Name:     "ONTAP Storage Virtual Machines",
		},
		{
			Factory:  dataSourceOpenzfsSnapshot,
			TypeName: "aws_fsx_openzfs_snapshot",
			Name:     "OpenZFS Snapshot",
		},
		{
			Factory:  dataSourceWindowsFileSystem,
			TypeName: "aws_fsx_windows_file_system",
			Name:     "Windows File System",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceBackup,
			TypeName: "aws_fsx_backup",
			Name:     "Backup",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDataRepositoryAssociation,
			TypeName: "aws_fsx_data_repository_association",
			Name:     "Data Repository Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceFileCache,
			TypeName: "aws_fsx_file_cache",
			Name:     "File Cache",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceLustreFileSystem,
			TypeName: "aws_fsx_lustre_file_system",
			Name:     "Lustre File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceONTAPFileSystem,
			TypeName: "aws_fsx_ontap_file_system",
			Name:     "ONTAP File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceONTAPStorageVirtualMachine,
			TypeName: "aws_fsx_ontap_storage_virtual_machine",
			Name:     "ONTAP Storage Virtual Machine",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceONTAPVolume,
			TypeName: "aws_fsx_ontap_volume",
			Name:     "ONTAP Volume",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceOpenZFSFileSystem,
			TypeName: "aws_fsx_openzfs_file_system",
			Name:     "OpenZFS File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceOpenZFSSnapshot,
			TypeName: "aws_fsx_openzfs_snapshot",
			Name:     "OpenZFS Snapshot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceOpenZFSVolume,
			TypeName: "aws_fsx_openzfs_volume",
			Name:     "OpenZFS Volume",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceWindowsFileSystem,
			TypeName: "aws_fsx_windows_file_system",
			Name:     "Windows File System",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.FSx
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*fsx.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*fsx.Options){
		fsx.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *fsx.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "fsx",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return fsx.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*fsx.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*fsx.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *fsx.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*fsx.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}

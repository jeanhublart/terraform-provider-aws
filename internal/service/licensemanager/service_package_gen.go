// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
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
			Factory:  dataSourceGrants,
			TypeName: "aws_licensemanager_grants",
			Name:     "Grants",
		},
		{
			Factory:  dataSourceReceivedLicense,
			TypeName: "aws_licensemanager_received_license",
			Name:     "Received License",
		},
		{
			Factory:  dataSourceReceivedLicenses,
			TypeName: "aws_licensemanager_received_licenses",
			Name:     "Received Licenses",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAssociation,
			TypeName: "aws_licensemanager_association",
			Name:     "Association",
		},
		{
			Factory:  resourceGrant,
			TypeName: "aws_licensemanager_grant",
			Name:     "Grant",
		},
		{
			Factory:  resourceGrantAccepter,
			TypeName: "aws_licensemanager_grant_accepter",
			Name:     "Grant Accepter",
		},
		{
			Factory:  resourceLicenseConfiguration,
			TypeName: "aws_licensemanager_license_configuration",
			Name:     "License Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.LicenseManager
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*licensemanager.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*licensemanager.Options){
		licensemanager.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *licensemanager.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "licensemanager",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return licensemanager.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*licensemanager.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*licensemanager.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *licensemanager.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*licensemanager.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}

// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
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
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newSourceAPIAssociationResource,
			TypeName: "aws_appsync_source_api_association",
			Name:     "Source API Association",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAPICache,
			TypeName: "aws_appsync_api_cache",
			Name:     "API Cache",
		},
		{
			Factory:  resourceAPIKey,
			TypeName: "aws_appsync_api_key",
			Name:     "API Key",
		},
		{
			Factory:  resourceDataSource,
			TypeName: "aws_appsync_datasource",
			Name:     "Data Source",
		},
		{
			Factory:  resourceDomainName,
			TypeName: "aws_appsync_domain_name",
			Name:     "Domain Name",
		},
		{
			Factory:  resourceDomainNameAPIAssociation,
			TypeName: "aws_appsync_domain_name_api_association",
			Name:     "Domain Name API Association",
		},
		{
			Factory:  resourceFunction,
			TypeName: "aws_appsync_function",
			Name:     "Function",
		},
		{
			Factory:  resourceGraphQLAPI,
			TypeName: "aws_appsync_graphql_api",
			Name:     "GraphQL API",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceResolver,
			TypeName: "aws_appsync_resolver",
			Name:     "Resolver",
		},
		{
			Factory:  resourceType,
			TypeName: "aws_appsync_type",
			Name:     "Type",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppSync
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*appsync.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*appsync.Options){
		appsync.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *appsync.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "appsync",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return appsync.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*appsync.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*appsync.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *appsync.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*appsync.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}

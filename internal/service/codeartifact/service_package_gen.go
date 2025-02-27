// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package codeartifact

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
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
			Factory:  dataSourceAuthorizationToken,
			TypeName: "aws_codeartifact_authorization_token",
			Name:     "Authoiration Token",
		},
		{
			Factory:  dataSourceRepositoryEndpoint,
			TypeName: "aws_codeartifact_repository_endpoint",
			Name:     "Repository Endpoint",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceDomain,
			TypeName: "aws_codeartifact_domain",
			Name:     "Domain",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceDomainPermissionsPolicy,
			TypeName: "aws_codeartifact_domain_permissions_policy",
			Name:     "Domain Permissions Policy",
		},
		{
			Factory:  resourceRepository,
			TypeName: "aws_codeartifact_repository",
			Name:     "Repository",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceRepositoryPermissionsPolicy,
			TypeName: "aws_codeartifact_repository_permissions_policy",
			Name:     "Repository Permissions Policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CodeArtifact
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*codeartifact.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*codeartifact.Options){
		codeartifact.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *codeartifact.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "codeartifact",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return codeartifact.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*codeartifact.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*codeartifact.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *codeartifact.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*codeartifact.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}

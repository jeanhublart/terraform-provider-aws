// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) EphemeralResources(ctx context.Context) []*types.ServicePackageEphemeralResource {
	return []*types.ServicePackageEphemeralResource{
		{
			Factory:  newEphemeralClusterAuth,
			TypeName: "aws_eks_cluster_auth",
			Name:     "ClusterAuth",
		},
	}
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory:  newPodIdentityAssociationResource,
			TypeName: "aws_eks_pod_identity_association",
			Name:     "Pod Identity Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "association_arn",
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceAccessEntry,
			TypeName: "aws_eks_access_entry",
			Name:     "Access Entry",
		},
		{
			Factory:  dataSourceAddon,
			TypeName: "aws_eks_addon",
			Name:     "Add-On",
		},
		{
			Factory:  dataSourceAddonVersion,
			TypeName: "aws_eks_addon_version",
			Name:     "Add-On Version",
		},
		{
			Factory:  dataSourceCluster,
			TypeName: "aws_eks_cluster",
			Name:     "Cluster",
		},
		{
			Factory:  dataSourceClusterAuth,
			TypeName: "aws_eks_cluster_auth",
			Name:     "Cluster Authentication Token",
		},
		{
			Factory:  dataSourceClusters,
			TypeName: "aws_eks_clusters",
			Name:     "Clusters",
		},
		{
			Factory:  dataSourceNodeGroup,
			TypeName: "aws_eks_node_group",
			Name:     "Node Group",
		},
		{
			Factory:  dataSourceNodeGroups,
			TypeName: "aws_eks_node_groups",
			Name:     "Node Groups",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccessEntry,
			TypeName: "aws_eks_access_entry",
			Name:     "Access Entry",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "access_entry_arn",
			},
		},
		{
			Factory:  resourceAccessPolicyAssociation,
			TypeName: "aws_eks_access_policy_association",
			Name:     "Access Policy Association",
		},
		{
			Factory:  resourceAddon,
			TypeName: "aws_eks_addon",
			Name:     "Add-On",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceCluster,
			TypeName: "aws_eks_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceFargateProfile,
			TypeName: "aws_eks_fargate_profile",
			Name:     "Fargate Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceIdentityProviderConfig,
			TypeName: "aws_eks_identity_provider_config",
			Name:     "Identity Provider Config",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceNodeGroup,
			TypeName: "aws_eks_node_group",
			Name:     "Node Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.EKS
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*eks.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*eks.Options){
		eks.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		func(o *eks.Options) {
			if region := config["region"].(string); o.Region != region {
				tflog.Info(ctx, "overriding provider-configured AWS API region", map[string]any{
					"service":         "eks",
					"original_region": o.Region,
					"override_region": region,
				})
				o.Region = region
			}
		},
		withExtraOptions(ctx, p, config),
	}

	return eks.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*eks.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*eks.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *eks.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*eks.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}

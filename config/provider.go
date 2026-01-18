/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	applicationkeyCluster "github.com/eaglesemanation/provider-b2/config/cluster/applicationkey"
	bucketCluster "github.com/eaglesemanation/provider-b2/config/cluster/bucket"
	bucketfileversionCluster "github.com/eaglesemanation/provider-b2/config/cluster/bucketfileversion"
	bucketnotificationrulesCluster "github.com/eaglesemanation/provider-b2/config/cluster/bucketnotificationrules"
	applicationkeyNamespaced "github.com/eaglesemanation/provider-b2/config/namespaced/applicationkey"
	bucketNamespaced "github.com/eaglesemanation/provider-b2/config/namespaced/bucket"
	bucketfileversionNamespaced "github.com/eaglesemanation/provider-b2/config/namespaced/bucketfileversion"
	bucketnotificationrulesNamespaced "github.com/eaglesemanation/provider-b2/config/namespaced/bucketnotificationrules"
)

const (
	resourcePrefix = "b2"
	modulePath     = "github.com/eaglesemanation/provider-b2"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		applicationkeyCluster.Configure,
		bucketCluster.Configure,
		bucketfileversionCluster.Configure,
		bucketnotificationrulesCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		applicationkeyNamespaced.Configure,
		bucketNamespaced.Configure,
		bucketfileversionNamespaced.Configure,
		bucketnotificationrulesNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

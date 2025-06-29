/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/eaglesemanation/provider-b2/config/application_key"
	"github.com/eaglesemanation/provider-b2/config/bucket"
	"github.com/eaglesemanation/provider-b2/config/bucket_file_version"
	"github.com/eaglesemanation/provider-b2/config/bucket_notification_rules"
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
		application_key.Configure,
		bucket.Configure,
		bucket_file_version.Configure,
		bucket_notification_rules.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/eaglesemanation/provider-b2/config/applicationkey"
	"github.com/eaglesemanation/provider-b2/config/bucket"
	"github.com/eaglesemanation/provider-b2/config/bucketfileversion"
	"github.com/eaglesemanation/provider-b2/config/bucketnotificationrules"
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
		applicationkey.Configure,
		bucket.Configure,
		bucketfileversion.Configure,
		bucketnotificationrules.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

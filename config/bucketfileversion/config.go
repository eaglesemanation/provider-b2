package bucketfileversion

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("b2_bucket_file_version", func(r *config.Resource) {
		r.ShortGroup = "b2"
		r.Kind = "BucketFileVersion"
		r.References["bucket_id"] = config.Reference{
			TerraformName:     "b2_bucket",
			RefFieldName:      "BucketRef",
			SelectorFieldName: "BucketSelector",
		}
	})
}

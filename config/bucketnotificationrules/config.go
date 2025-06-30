package bucketnotificationrules

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("b2_bucket_notification_rules", func(r *config.Resource) {
		r.ShortGroup = "b2"
		r.Kind = "BucketNotificationRules"
		r.References["bucket_id"] = config.Reference{
			TerraformName:     "b2_bucket",
			RefFieldName:      "BucketRef",
			SelectorFieldName: "BucketSelector",
		}
	})
}

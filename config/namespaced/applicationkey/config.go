package applicationkey

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("b2_application_key", func(r *config.Resource) {
		r.ShortGroup = "b2"
		r.Kind = "ApplicationKey"
		r.ExternalName.OmittedFields = []string{
			"bucket_id",
		}
		r.References["bucket_ids"] = config.Reference{
			TerraformName: "b2_bucket",
		}
		if s, ok := r.TerraformResource.Schema["application_key_id"]; ok {
			s.Sensitive = true
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["application_key_id"].(string); ok {
				conn["application_key_id"] = []byte(a)
			}
			if a, ok := attr["application_key"].(string); ok {
				conn["application_key"] = []byte(a)
			}
			return conn, nil
		}
	})
}

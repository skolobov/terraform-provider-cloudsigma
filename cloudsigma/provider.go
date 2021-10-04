package cloudsigma

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a schema.Provider for CloudSigma.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDSIGMA_USERNAME", nil),
				Description: "The CloudSigma user email.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDSIGMA_PASSWORD", nil),
				Description: "The CloudSigma password.",
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDSIGMA_LOCATION", "zrh"),
				Description: "The location endpoint for CloudSigma. Default is 'zrh'.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"cloudsigma_ip":            dataSourceCloudSigmaIP(),
			"cloudsigma_library_drive": dataSourceCloudSigmaLibraryDrive(),
			"cloudsigma_drive":         dataSourceCloudSigmaDrive(),
			"cloudsigma_license":       dataSourceCloudSigmaLicense(),
			"cloudsigma_location":      dataSourceCloudSigmaLocation(),
			"cloudsigma_profile":       dataSourceCloudSigmaProfile(),
			"cloudsigma_subscription":  dataSourceCloudSigmaSubscription(),
			"cloudsigma_tag":           dataSourceCloudSigmaTag(),
			"cloudsigma_vlan":          dataSourceCloudSigmaVLAN(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"cloudsigma_acl":             resourceCloudSigmaACL(),
			"cloudsigma_drive":           resourceCloudSigmaDrive(),
			"cloudsigma_firewall_policy": resourceCloudSigmaFirewallPolicy(),
			"cloudsigma_remote_snapshot": resourceCloudSigmaRemoteSnapshot(),
			"cloudsigma_server":          resourceCloudSigmaServer(),
			"cloudsigma_snapshot":        resourceCloudSigmaSnapshot(),
			"cloudsigma_ssh_key":         resourceCloudSigmaSSHKey(),
			"cloudsigma_tag":             resourceCloudSigmaTag(),
		},
	}

	provider.ConfigureContextFunc = providerConfigure(provider)

	return provider
}

func providerConfigure(provider *schema.Provider) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		config := &Config{
			Username: d.Get("username").(string),
			Password: d.Get("password").(string),
			Location: d.Get("location").(string),
		}

		config.loadAndValidate(ctx, provider.TerraformVersion)

		return config.Client(), nil
	}
}

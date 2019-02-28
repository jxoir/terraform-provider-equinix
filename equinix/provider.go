package equinix

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns terraform.ResourceProvider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"equinix_app_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is the Equinix Developer Applicaition ID",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"EQUINIX_API_ID"}, nil),
			},
			"equinix_app_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is the Equinix Developer Applicaition Secret",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"EQUINIX_API_SECRET"}, nil),
			},
			"equinix_api_user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is the Equinix Customer Portal user to make API calls",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"ECX_API_USER"}, nil),
			},
			"equinix_api_password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is the Equinix Customer Portal user password to make API calls",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"ECX_API_USER_PASSWORD"}, nil),
			},
			"equinix_api_host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is the Equinix Customer Portal user password to make API calls",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"ECX_API_HOST"}, nil),
			},
			"debug": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"equinix_ecx_connection": resourceConnection(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AppID:     d.Get("equinix_app_id").(string),
		AppSecret: d.Get("equinix_app_secret").(string),
		UserName:  d.Get("equinix_api_user").(string),
		Password:  d.Get("equinix_api_password").(string),
		Endpoint:  d.Get("equinix_api_host").(string),
		Debug:     d.Get("debug").(bool),
	}

	return config.Client()
}

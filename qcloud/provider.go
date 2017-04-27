package qcloud

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"secret_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				//			DefaultFunc: schema.EnvDefaultFunc("QCLOUD_SECRET_ID", ""),
				Description: descriptions["secret_id"],
			},

			"secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				//			DefaultFunc: schema.EnvDefaultFunc("QCLOUD_SECRET_KEY", ""),
				Description: descriptions["secret_key"],
			},

			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				//			DefaultFunc: schema.EnvDefaultFunc("QCLOUD_REGION", "gz"),
				Description: descriptions["region"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"qcloud_clb": resourceQcloudClb(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (meta interface{}, err error) {
	return &Client{
		Region:    d.Get("region").(string),
		SecretId:  d.Get("secret_id").(string),
		SecretKey: d.Get("secret_key").(string),
	}, nil
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"region":     "Qcloud region to use for API requests. Defaults to \"gz\" if blank.",
		"secret_id":  "Qcloud SecretId to use for API requests. Defaults to \"guest\" if blank.",
		"secret_key": "Qcloud SecretKey to use for API requests. Defaults to \"guest\" if blank.",
	}
}

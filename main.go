package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/tpounds/terraform-provider-qcloud/qcloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: qcloud.Provider,
	})
}

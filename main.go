package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	// "github.com/awesome/terraform-provider-pingdom/pingdom"
	"github.com/awesome-ping/terraform-provider-awesome-ping/pingdom"
)

func main() {
	providerserver.Serve(context.Background(), pingdom.New, providerserver.ServeOpts{
		// Address: "registry.terraform.io/mbarper/pingdom",
		Address: "registry.terraform.io/awesome-ping/pingdom",
	})
}

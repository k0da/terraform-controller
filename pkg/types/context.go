package types

import (

	tfv1 "github.com/rancher/terraform-controller/pkg/generated/controllers/terraformcontroller.cattle.io"
	core "github.com/rancher/wrangler-api/pkg/generated/controllers/core"
)

type Context struct {
	Namespace string
	Tfv1   *tfv1.Factory
	Core   *core.Factory
}
func NewContext(namespace string, tfFactory *tfv1.Factory, coreFactory *core.Factory) *Context {
	context := &Context{
		Namespace: namespace,
			Tfv1:      tfFactory,
			Core:      coreFactory,
	}
	return context
}

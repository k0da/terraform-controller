package types

import (

	tfv1 "github.com/rancher/terraform-controller/pkg/generated/controllers/terraformcontroller.cattle.io"
	core "github.com/rancher/wrangler/pkg/generated/controllers/core"
)

type Context struct {
	Tfv1   *tfv1.Factory
	Core   *core.Factory
}
func NewContext(tfFactory *tfv1.Factory, coreFactory *core.Factory) *Context {
	context := &Context{
		Tfv1:      tfFactory,
		Core:      coreFactory,
	}
	return context
}

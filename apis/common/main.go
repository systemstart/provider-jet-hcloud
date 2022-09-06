package common

import (
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	tjresource "github.com/crossplane/terrajet/pkg/resource"
)

func ExtractResourceID() reference.ExtractValueFn {
	return func(mr resource.Managed) string {
		tr, ok := mr.(tjresource.Terraformed)
		if !ok {
			return ""
		}
		return tr.GetID()
	}
}

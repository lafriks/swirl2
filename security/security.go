package security

import (
	"github.com/cuigh/auxo/app/ioc"
)

const PkgName = "security"

func init() {
	ioc.Put(NewIdentifier, ioc.Name("identifier"))
	ioc.Put(NewAuthorizer, ioc.Name("authorizer"))
}

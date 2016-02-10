package security

import (
	"github.com/mariusmagureanu/vaed/response"
)

type AclPlugin struct {

}

func (AclPlugin) FilterBody() {

}

func (AclPlugin) FilterHeaders() {

}

func (AclPlugin) Log() {

}

func (AclPlugin) Run(out *response.Out) {
	out.Headers["X-Acl"] = "1000"
}

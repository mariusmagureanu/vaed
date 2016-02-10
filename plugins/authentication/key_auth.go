package authentication

import (
	"github.com/mariusmagureanu/vaed/response"
)

type KeyAuthPlugin struct {

}

func (KeyAuthPlugin) FilterBody() {

}

func (KeyAuthPlugin) FilterHeaders() {

}

func (KeyAuthPlugin) Log() {

}

func (KeyAuthPlugin) Run(out *response.Out) {
	out.Headers["X-Key-Auth"] = "1000"
}

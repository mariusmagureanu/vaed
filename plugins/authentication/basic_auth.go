package authentication

import (
	"vaed/response"
)

type BasicAuthPlugin struct {

}

func (BasicAuthPlugin) FilterBody() {

}

func (BasicAuthPlugin) FilterHeaders() {

}

func (BasicAuthPlugin) Log() {

}

func (BasicAuthPlugin) Run(out *response.Out) {
	out.Headers["X-Basic-Auth"] = "1000"
}

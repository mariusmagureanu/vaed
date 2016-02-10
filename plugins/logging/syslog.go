package logging

import (
	"github.com/mariusmagureanu/vaed/response"
)

type SysLogPlugin struct {

}

func (SysLogPlugin) FilterBody() {

}

func (SysLogPlugin) FilterHeaders() {

}

func (SysLogPlugin) Log() {

}

func (SysLogPlugin) Run(out *response.Out) {
	out.Headers["X-Syslog"] = "1000"
}

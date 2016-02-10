package throttle

import (
	"github.com/mariusmagureanu/vaed/response"
)

type RateLimitPlugin struct {

}

func (RateLimitPlugin) FilterBody() {

}

func (RateLimitPlugin) FilterHeaders() {

}

func (RateLimitPlugin) Log() {

}

func (RateLimitPlugin) Run(out *response.Out) {
	out.Headers["X-Rate-Limit"] = "1000"
}

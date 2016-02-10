package metrics

import (
	"vaed/response"
)

type MetricsPlugin struct {

}

func (MetricsPlugin) FilterBody() {

}

func (MetricsPlugin) FilterHeaders() {

}

func (MetricsPlugin) Log() {

}

func (MetricsPlugin) Run(out *response.Out) {
	out.Headers["X-Metrics"] = "1000"
}

package plugins

import (
	"errors"
	"fmt"
	"vaed/plugins/authentication"
	"vaed/plugins/logging"
	"vaed/plugins/metrics"
	"vaed/plugins/security"
	"vaed/plugins/throttle"
	"vaed/response"
)

func init() {
	Register("basic-auth", func() BasePlugin {
		return &authentication.BasicAuthPlugin{}
	})
	Register("key-auth", func() BasePlugin {
		return &authentication.KeyAuthPlugin{}
	})
	Register("syslog", func() BasePlugin {
		return &logging.SysLogPlugin{}
	})

	Register("metrics", func() BasePlugin {
		return &metrics.MetricsPlugin{}
	})

	Register("acl", func() BasePlugin {
		return &security.AclPlugin{}
	})

	Register("rate-limit", func() BasePlugin {
		return &throttle.RateLimitPlugin{}
	})

}

type BasePlugin interface {
	FilterBody()
	FilterHeaders()
	Log()
	Run(out *response.Out)
}

type Factory func() BasePlugin
var plugins = make(map[string]Factory)

func Register(name string, f Factory) error {
	if _, exists := plugins[name]; exists {
		return fmt.Errorf("Duplicated plugin!", name)
	}
	plugins[name] = f
	return nil
}

//Get gets plugins
func Get(name string) (BasePlugin, error) {
	factory := plugins[name]
	if factory == nil {
		return nil, errors.New("Plugin not found.")
	}
	return factory(), nil
}

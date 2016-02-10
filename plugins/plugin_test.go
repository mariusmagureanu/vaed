package plugins

import (
	"testing"
	"vaed/plugins/authentication"
	"vaed/plugins/logging"
	"vaed/plugins/metrics"
	"vaed/plugins/security"
	"vaed/plugins/throttle"
)

func TestRegisterBasicAuth(t *testing.T) {
	err := Register("p1", func() BasePlugin {
		return &authentication.BasicAuthPlugin{}
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterKeyAuth(t *testing.T) {
	err := Register("p2", func() BasePlugin {
		return &authentication.KeyAuthPlugin{}
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterSysLog(t *testing.T) {
	err := Register("p3", func() BasePlugin {
		return &logging.SysLogPlugin{}
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterMetrics(t *testing.T) {
	err := Register("p4", func() BasePlugin {
		return &metrics.MetricsPlugin{}
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterAcl(t *testing.T) {
	err := Register("p5", func() BasePlugin {
		return &security.AclPlugin{}
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterRateLimit(t *testing.T) {
	err := Register("p6", func() BasePlugin {
		return &throttle.RateLimitPlugin{}
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterDuplicated(t *testing.T) {
	err := Register("p1", func() BasePlugin {
		return &authentication.KeyAuthPlugin{}
	})
	if err == nil {
		t.Fatal("expected to fail duplicated registration")
	}
}

func TestGet(t *testing.T) {
	ep1, _ := Get("p1")

	if _, ok:=ep1.(*authentication.BasicAuthPlugin); !ok {
		t.Fatal("plugin returned unexpected instance")
	}
}

func TestGetUnknown(t *testing.T) {
	_, err := Get("unknown")
	if err == nil {
		t.Fatal("should have blown here, no plugin found by that name.")
	}
}

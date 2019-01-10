package ctlcmd

import (
	"testing"
	"v2ray.com/v2ray-study/common/platform"
)

func TestGetToolLocation(t *testing.T) {
	v2ctl := platform.GetToolLocation("v2ctl")
	t.Log(v2ctl)
}

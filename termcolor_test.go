package termcolor

import (
	"testing"
)

func TestCapability(t *testing.T) {
	caps, _ := Capability()
	//Should be ColorNone, Color8, Color16, Color256, ColorTrue
	t.Logf("Capability: %d", caps)
}

func TestCapability_ColorSupport(t *testing.T) {
	caps, status := Capability()
	if status != Uncertain && caps == ColorNone {
		t.Log("Terminal doesn't support color")
	}
}

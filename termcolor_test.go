package termcolor

import (
	"os"
	"testing"
)

func TestCapability(t *testing.T) {
	caps := Capability()
	//Should be ColorNone, Color8, Color16, Color256, ColorTrue
	t.Logf("Capability: %d", caps)
}

func TestDisabled(t *testing.T) {
	os.Setenv("NO_COLOR", "1")
	if !IsDisabled() {
		t.Errorf("Disabled() should return true when NO_COLOR is set")
	}

	os.Unsetenv("NO_COLOR")
	if IsDisabled() {
		t.Errorf("Disabled() should return false when NO_COLOR is not set")
	}
}

func TestIsTerminal(t *testing.T) {
	t.Logf("Is Terminal: %v", IsTerminal())
}



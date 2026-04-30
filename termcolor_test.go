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


func TestIsTerminal(t *testing.T) {
	t.Logf("Is Terminal: %v", IsTerminal())
}


func TestDisabled(t *testing.T) {
	os.Setenv("NO_COLOR", "1")
	if !Disabled() {
		t.Errorf("Disabled() should return true when NO_COLOR is set")
	}

	os.Unsetenv("NO_COLOR")
	if Disabled() {
		t.Errorf("Disabled() should return false when NO_COLOR is not set")
	}
}


func TestEnabled(t *testing.T) {
	os.Setenv("FORCE_COLOR", "1")
	if !Enabled() {
		t.Errorf("Enabled() should return true when FORCE_COLOR is set")
	}

	os.Unsetenv("FORCE_COLOR")
	if Enabled() {
		t.Errorf("Enabled() should return false when FORCE_COLOR is not set")
	}
}

func TestSupportsColor(t *testing.T) {
	os.Setenv("NO_COLOR", "1")
	if SupportsColor() {
		t.Log("SUPPP: ", SupportsColor())
		t.Errorf("SupportsColor() should return false when NO_COLOR is set")
	}

	os.Setenv("FORCE_COLOR", "1")
	if !SupportsColor() {
		t.Errorf("SupportsColor() should return true when FORCE_COLOR is set")
	}

	os.Unsetenv("NO_COLOR")
	t.Logf("SupportsColor: %v", SupportsColor())
}
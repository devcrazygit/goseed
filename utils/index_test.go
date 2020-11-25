package utils

import (
	"os"
	"testing"
)

func TestDefaultVal(t *testing.T) {
	val := EnvVar("FOO", "BAR")
	if val != "BAR" {
		t.Errorf("Expected BAR but got %s", val)
	}
}

func TestEnvVal(t *testing.T) {
	os.Setenv("FOO", "BAR2")
	val := EnvVar("FOO", "BAR")
	if val != "BAR2" {
		t.Errorf("Expected BAR2 but got %s", val)
	}
}

package lib

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	var testEnv = struct {
		key   string
		value string
	}{"GO_ENV", "TEST"}
	os.Setenv(testEnv.key, testEnv.value)

	env := GetEnv(testEnv.key, "default")

	if env != testEnv.value {
		t.Fatalf("got %s wanted %s", env, testEnv.value)
	}

	defaulted := GetEnv("RANDOM_ENV_VARIABLE", "default")

	if defaulted != "default" {
		t.Fatalf("got %s wanted %s", defaulted, "default")
	}

}

package loader_test

import (
	"testing"

	"github.com/adamcohen/godotenv-test/loader"
)

func TestLoadVars(t *testing.T) {
	res := loader.LoadVars()

	if res != "MY_VALUE" {
		t.Fatalf("test failed")
	}
}

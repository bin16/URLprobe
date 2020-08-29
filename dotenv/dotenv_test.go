package dotenv

import (
	"os"
	"testing"
)

func TestDotEnv(t *testing.T) {
	if v := os.Getenv("TEST"); v != "2020" {
		t.Errorf("Failed to get env <TEST>, got <%s>, want <%s>", v, "2020")
	}
}

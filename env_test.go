package env

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestGet(t *testing.T) {
	g := Get("KEY1", Default("VAL1"))
	if g != "VAL1" {
		t.Errorf("expected %s, got %s (key %s)", "VAL1", g, "KEY1")
	}
}

func TestSet(t *testing.T) {
	Set("KEY2", "VALUE2")
	g := Get("KEY2")
	if g != "VALUE2" {
		t.Errorf("expected %s, got %s (key %s)", "VALUE2", g, "KEY2")
	}
}

func TestNotDefault(t *testing.T) {
	g := Get("GOPATH", Default("/home/your-user/go"))
	if g == "/home/your-user/go" {
		t.Errorf("expected not to get default value")
	}
}

func TestTransform(t *testing.T) {
	fn := func(value string, found bool) string {
		return "THIS!"
	}
	g := Get("SOMEKEY", Transform(fn))
	if g != "THIS!" {
		t.Errorf("should have applied transform function")
	}
}

func ExampleEnvVars() {
	Set("THIS", "TO-THAT")
	fmt.Println(Get("THIS"))
	// Output: TO-THAT
}

func ExampleDefault() {
	fmt.Println(Get("THAT OVER THERE", Default("OH LOOK IS THIS")))
	// Output: OH LOOK IS THIS
}

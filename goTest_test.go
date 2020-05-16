package gotest


import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	if got != "Hello v1" {
		t.Errorf("Hello function not working")
	}
}

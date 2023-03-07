package go_hello_world

import "testing"

func TestGreeter(t *testing.T) {
	expected := "Hello world!"
	actual := HelloWorld()
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
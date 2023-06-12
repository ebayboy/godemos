package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	if actual := HelloWorld(); actual != expected {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestHelloWorld2(t *testing.T) {
	expected := "Hello, World2!"
	if actual := HelloWorld2(); actual != expected {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

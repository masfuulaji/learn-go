package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Dani")

	if result != "Hello Dani" {
		t.Error("Result must be 'Hello Dani'")
	}
}

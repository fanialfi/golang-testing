package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("fani")

	if result != "Hello fani" {
		// error
		panic(`result must be "Hello fani"`)
	}
}

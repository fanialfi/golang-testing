package helper

import (
	"fmt"
	"log"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("tidak bisa menjalankan test di windows")
	}

	expected := "Hello fani"
	result := HelloWorld("fani")
	require.Equal(t, expected, result, "result must be \"", expected, "\"")
	log.Println("TestHelloWorld with require done")
}

func TestHelloWorldRequire(t *testing.T) {
	expected := "Hello fani "
	result := HelloWorld("fani")
	require.Equal(t, expected, result, "result must be \"", expected, "\"")
	log.Println("TestHelloWorld with require done")
}

func TestHelloWorldAssert(t *testing.T) {
	expected := "Hello fani "
	result := HelloWorld("fani")
	assert.Equal(t, expected, result, "result must be \"", expected, "\"")
	log.Println("TestHelloWorld with assert done")
}

func TestHelloWorld(t *testing.T) {
	respect := "Hello fani "
	result := HelloWorld("fani")

	if result != respect {
		// menggagalkan unit test menggunakan panic bukanlah hal yang bagus
		// panic(`result must be "Hello fani"`)

		// t.Fail() // karena kita tidak tahu nanti penyebab fail nya apa
		t.Errorf(`testing failed because result not "%s" result is : "%s"`, respect, result)
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldKedua(t *testing.T) {
	respect := "Hello fani "
	result := HelloWorld("fani")

	if result != respect {
		t.Fatalf(`testing failed because result not "%s" result is : "%s"`, respect, result)
	}

	fmt.Println("TestHelloWorldKedua Done")
}

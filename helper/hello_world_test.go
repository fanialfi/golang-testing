package helper

import (
	"fmt"
	"log"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("fanialfi")
	}
}

// sub test yang lebih dinamis
func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "fani",
			request:  "fani",
			expected: "Hello fani",
		},
		{
			name:     "alfi",
			request:  "alfi",
			expected: "Hello alfi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("SubTestOne", func(t *testing.T) {
		expected := "Hello fanialfi"
		result := HelloWorld("fanialfi")
		require.Equalf(t, expected, result, "hasil yang di harapkan adalah \"%s\"", expected)
	})

	t.Run("SubTestTwo", func(t *testing.T) {
		expected := "Hello fanialfi"
		result := HelloWorld("fanialfi")
		require.Equalf(t, expected, result, "hasil yang di harapkan adalah \"%s\"", expected)
	})
}

func TestMain(m *testing.M) {
	// before
	log.Println("BEFORE UNIT TEST")
	m.Run()
	// after
	log.Println("AFTER UNIT TEST")
}

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
	expected := "Hello fani"
	result := HelloWorld("fani")
	require.Equal(t, expected, result, "result must be \"", expected, "\"")
	log.Println("TestHelloWorld with require done")
}

func TestHelloWorldAssert(t *testing.T) {
	expected := "Hello fani"
	result := HelloWorld("fani")
	assert.Equal(t, expected, result, "result must be \"", expected, "\"")
	log.Println("TestHelloWorld with assert done")
}

func TestHelloWorld(t *testing.T) {
	respect := "Hello fani"
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
	respect := "Hello fani"
	result := HelloWorld("fani")

	if result != respect {
		t.Fatalf(`testing failed because result not "%s" result is : "%s"`, respect, result)
	}

	fmt.Println("TestHelloWorldKedua Done")
}

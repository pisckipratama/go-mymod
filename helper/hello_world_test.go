package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TESTING")
	m.Run()
	fmt.Println("AFTER UNIT TESTING")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}

	result := HelloWorld("Piscki")
	require.Equal(t, "Hello Piscki", result, "Result must be 'Hello Piscki'")
	fmt.Println("testing require done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Piscki")

	require.Equal(t, "Hello Piscki", result, "Result must be 'Hello Piscki'")
	fmt.Println("testing require done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Piscki")

	assert.Equal(t, "Hello Piscki", result, "Result must be 'Hello Piscki'")
	fmt.Println("testing assertion done")
}

func TestHelloWorldPiscki(t *testing.T) {
	result := HelloWorld("Piscki")

	if result != "Hello Piscki" {
		t.Error("Result must be 'Hello Piscki'")
	}

	fmt.Println("TestHelloWorldPiscki Done")
}

func TestHelloWorldPratama(t *testing.T) {
	result := HelloWorld("Pratama")

	if result != "Hello Pratama" {
		t.Fatal("Result must be 'Hello Pratama'")
	}

	fmt.Println("TestHelloWorldPratama Done")
}

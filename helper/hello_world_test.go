package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Piscki")

	require.Equal(t, "Hi Piscki", result, "Result must be 'Hi Piscki'")
	fmt.Println("testing require done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Piscki")

	assert.Equal(t, "Hi Piscki", result, "Result must be 'Hi Piscki'")
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

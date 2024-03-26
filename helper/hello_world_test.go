package helper

import (
	"fmt"
	"testing"
)

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

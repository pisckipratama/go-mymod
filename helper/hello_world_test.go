package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Piscki")

	if result != "Hello Piscki" {
		panic("Test is not Hello Piscki")
	}
}

func TestHelloWorldPratama(t *testing.T) {
	result := HelloWorld("Pratama")

	if result != "Hello Pratama" {
		panic("Test is not Hello Pratama")
	}
}

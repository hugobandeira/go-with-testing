package main

import "testing"

func TestHelloSuccess(testing *testing.T) {
	result := Hello("Hugo")
	expected := "Fist Testing Hugo Hugo"

	if result != expected {
		testing.Errorf("incorrect value expected '%s' value '%s'", expected, result)
	}
}

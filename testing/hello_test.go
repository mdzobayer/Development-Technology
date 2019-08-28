// https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
package main

import "testing"

func TestHello(t *testing.T) {
	emptyResult := hello("") // should return "Hello Dude!"

	if emptyResult != "Hello Dude!" {
		t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Dude!", emptyResult)
	} else {
		t.Logf("hello(\"\") success, expected %v, got %v", "Hello Dude!", emptyResult)
	}

	result := hello("John") // should return "Hello John!"

	if result != "Hello John!" {
		t.Errorf("hello(\"John\") failed, expected %v, got %v", "Hello John!", result)
	} else {
		t.Logf("hello(\"John\") success, expected %v, got %v", "Hello John!", result)
	}
}

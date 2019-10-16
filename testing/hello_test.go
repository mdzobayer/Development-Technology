// https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
// Command: go test -v
// Command converfile: go test -coverprofile=cover.txt
// Command html: go tool cover -html=cover.txt -o cover.html
package main

import "testing"

// test hello function with empty argument
func TestHelloEmptyArg(t *testing.T) {

	emptyResult := hello("") // should return "Hello Dude!"

	if emptyResult != "Hello Dude!" {
		t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello Dude!", emptyResult)
	} else {
		t.Logf("hello(\"\") success, expected %v, got %v", "Hello Dude!", emptyResult)
	}
}

// test hello function with valid argument

func TestHelloValidArg(t *testing.T) {

	// test for valid  argument
	result := hello("John") // should return "Hello John!"

	if result != "Hello John!" {
		t.Errorf("hello(\"John\") failed, expected %v, got %v", "Hello John!", result)
	} else {
		t.Logf("hello(\"John\") success, expected %v, got %v", "Hello John!", result)
	}
}

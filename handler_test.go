package main

import "testing"

func TestHandle(t *testing.T) {
	if handle("blah") != "Hi there, I love blah!" {
		t.Errorf("Failed")
	}
}

package main

import "testing"

func TestHello(t *testing.T) {
	var got string = SayWhat("Turner")
	want := "Turner"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

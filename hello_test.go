package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Soubhik")
	want := "Hello Soubhik"

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}

}

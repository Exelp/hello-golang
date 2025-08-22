package main

import "testing"

func TestHello(t *testing.T) {
	want := "hello Go"
	got := hello()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

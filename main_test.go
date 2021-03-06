package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	in := "aaaaa"
	in2 := "a*"
	out := isMatch(in, in2)
	want := true
	if out != want {
		t.Errorf("got %t, want %t", out, want)
	}

	in = "áááááááááá"
	in2 = "á*"
	out = isMatch(in, in2)
	want = true
	if out != want {
		t.Errorf("got %t, want %t", out, want)
	}

	in = "ab"
	in2 = ".*"
	out = isMatch(in, in2)
	want = true
	if out != want {
		t.Errorf("got %t, want %t", out, want)
	}
}

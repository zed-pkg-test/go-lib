package golib

import "testing"

func TestGreetIdentifiesTheImmutableZedPackageNamespace(t *testing.T) {
	got := Greet("consumer")
	want := "hello consumer from zed-pkg-test/go-lib"
	if got != want {
		t.Fatalf("Greet(%q) = %q, want %q", "consumer", got, want)
	}
}

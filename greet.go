// Package golib is the one thing this zed-sourced module exports.
package golib

import "fmt"

// Greet names the immutable Zed package namespace so a consumer can prove which
// package it actually resolved, rather than merely that *something* compiled.
func Greet(who string) string {
	return fmt.Sprintf("hello %s from zed-pkg-test/go-lib", who)
}

# go-lib (zed-pkg-test/go-lib)

A trivial Go module published to the zed registry. Consumers source it **via
zed** (into their configured `[install].dir`) while the Go module proxy keeps
owning the rest of `go.mod`. See `.zpkg.toml`.

It also carries a normal `go.mod` so the Go toolchain resolves it as an ordinary
module once zed drops it into place — the consumer adds a `replace` pointing at
the install dir and never has to know a second package manager was involved.

# gofips

gofips is a linter that warns you about encryption usage that may not be FIPS-validated.

This linter uses the
[`go/analysis`](https://pkg.go.dev/golang.org/x/tools/go/analysis) API for
common linting functionality and conventions.

```console
$ go run cmd/gofips/main.go -- ./example.go
/home/lbragstad/go/src/github.com/rhmdnd/gofips/example.go:35:18: Seal is not a FIPS-validated implementation.
exit status 3
```

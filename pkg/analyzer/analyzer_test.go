package analyzer

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %s", err)
	}
	// Another common approach used by other linters is to put sample code
	// in a `testdata/` directory, then pass that directory location into
	// `Run(...)`. However, this forces `Run()` to load the packages for
	// those modules based on the `testdata/` directory being the GOPATH.
	// This causes issues when the example code we want to lint has a
	// dependency on something like x/crypto, but the `Run()` function
	// can't discover it (presumably due to opinionated usage of GOPATH?).
	// Working around this by keeping the example in the root of the
	// repository, so when the module is loaded it doesn't fail trying to
	// find golang.org/x/crypto.
	rootDir := filepath.Dir(filepath.Dir(cwd))
	analysistest.Run(t, rootDir, Analyzer, "example.go")
}

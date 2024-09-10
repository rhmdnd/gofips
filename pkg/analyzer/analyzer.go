package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "gofips",
	Doc:  "Checks for cryptographic usage from libraries not compliant with FIPS.",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		ce, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}
		// TODO: Add a module check to make sure we're scoping the
		// function check below to the correct module.

		// This could be generalized to detect any implementation that
		// hasn't been FIPS certified.
		if match := funcMatch(ce.Fun, "Seal"); !match {
			return true
		}
		funcName := getFuncName(ce.Fun)
		pass.Reportf(node.Pos(), "%s is not a FIPS-validated implementation.", funcName)
		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}

func getFuncName(e ast.Expr) string {
	return e.(*ast.SelectorExpr).Sel.Name
}

func funcMatch(e ast.Expr, m string) bool {
	sel, ok := e.(*ast.SelectorExpr)
	return ok && isIdent(sel.Sel, m)
}

func isIdent(e ast.Expr, i string) bool {
	id, ok := e.(*ast.Ident)
	return ok && id.Name == i
}

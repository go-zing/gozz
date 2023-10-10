package zutils

import (
	"go/ast"
)

// AssertFuncType to assert interface fields as function type and try return name
func AssertFuncType(field *ast.Field) (name string, ft *ast.FuncType, ok bool) {
	ft, ok = field.Type.(*ast.FuncType)
	if !ok || len(field.Names) == 0 {
		return
	}
	name = field.Names[0].Name
	return
}

// StructFields extracts struct fields names
func StructFields(typ *ast.StructType) (fields []string) {
	if typ.Fields == nil {
		return
	}

	anonymous := func(spec ast.Expr) (name *ast.Ident) {
		switch t := spec.(type) {
		case *ast.StarExpr:
			name, _ = t.X.(*ast.Ident)
		case *ast.SelectorExpr:
			name, _ = t.X.(*ast.Ident)
		case *ast.Ident:
			name = t
		}
		return
	}

	add := func(ident *ast.Ident) {
		if ident != nil && ident.IsExported() {
			fields = append(fields, ident.Name)
		}
	}

	for _, field := range typ.Fields.List {
		// anonymous field
		if len(field.Names) == 0 {
			add(anonymous(field.Type))
			continue
		}

		// with name
		for _, name := range field.Names {
			add(name)
		}
	}
	return
}

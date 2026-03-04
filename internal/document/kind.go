package document

import (
	"go/types"

	"github.com/sourcegraph/scip/bindings/go/scip"
)

// KindForObject maps a Go types.Object to its corresponding SCIP SymbolInformation_Kind.
func KindForObject(obj types.Object) scip.SymbolInformation_Kind {
	if obj == nil {
		return scip.SymbolInformation_UnspecifiedKind
	}

	switch v := obj.(type) {
	case *types.Func:
		sig, ok := v.Type().(*types.Signature)
		if ok && sig.Recv() != nil {
			return scip.SymbolInformation_Method
		}
		return scip.SymbolInformation_Function

	case *types.TypeName:
		if v.IsAlias() {
			return scip.SymbolInformation_TypeAlias
		}
		switch v.Type().Underlying().(type) {
		case *types.Struct:
			return scip.SymbolInformation_Struct
		case *types.Interface:
			return scip.SymbolInformation_Interface
		default:
			return scip.SymbolInformation_Type
		}

	case *types.Var:
		if v.IsField() {
			return scip.SymbolInformation_Field
		}
		return scip.SymbolInformation_Variable

	case *types.Const:
		return scip.SymbolInformation_Constant

	case *types.PkgName:
		return scip.SymbolInformation_Package

	case *types.Label:
		return scip.SymbolInformation_Variable

	default:
		return scip.SymbolInformation_UnspecifiedKind
	}
}

// Code generated by gozz:doc. DO NOT EDIT.

package doc

var (
	_types_doc = map[interface{}]map[string]string{
		(*T3)(nil): _doc_T3,

		(*T)(nil):  _doc_T,
		(*T2)(nil): _doc_T2,
		(*T4)(nil): _doc_T4,
		(*T5)(nil): _doc_T5,
		(*T6)(nil): _doc_T6,
		(*T7)(nil): _doc_T7,
	}

	_values_doc = map[string]map[interface{}]string{
		"variable": map[interface{}]string{
			ValueString:  "this is a string value",
			ValueInt:     "this is an int value",
			ValueStruct:  "this is a struct value",
			ValuePointer: "this is a pointer value",
			ValueInlineA: "they are inline declaration value",
			ValueInlineB: "they are inline declaration value",
			ValueInlineC: "they are inline type declaration value",
			ValueInlineD: "they are inline type declaration value",
		},
		"const": map[interface{}]string{
			ConstantString:  "this is a constant string",
			ConstantInt1:    "this is a constant int",
			ConstantInt2:    "this is another constant int",
			ConstantString2: "this is a single declared constant",
		},
	}
)

var _doc_T3 = map[string]string{
	"":       "this is an interface type",
	"Method": "this is an interface method",
}

var _doc_T = map[string]string{
	"":      "this is a struct type",
	"Field": "this is a struct field",
}

func (T) FieldDoc(f string) string { return _doc_T[f] }

var _doc_T2 = map[string]string{
	"":       "this is another struct type declared in group",
	"Field":  "this is a struct field\nthis is a struct field comment",
	"Field2": "this is another struct field",
}

func (T2) FieldDoc(f string) string { return _doc_T2[f] }

var _doc_T4 = map[string]string{
	"": "this is a refer type",
}

func (T4) FieldDoc(f string) string { return _doc_T4[f] }

var _doc_T5 = map[string]string{
	"": "this is a map type",
}

func (T5) FieldDoc(f string) string { return _doc_T5[f] }

var _doc_T6 = map[string]string{
	"": "this is an array type",
}

func (T6) FieldDoc(f string) string { return _doc_T6[f] }

var _doc_T7 = map[string]string{
	"": "this is another array type\nthis type has extra comments",
}

func (T7) FieldDoc(f string) string { return _doc_T7[f] }

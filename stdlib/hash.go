package stdlib

// Generated by 'goexports hash'. Do not edit!

import (
	"hash"
	"reflect"
)

func init() {
	Value["hash"] = map[string]reflect.Value{}
	Type["hash"] = map[string]reflect.Type{
		"Hash":   reflect.TypeOf((*hash.Hash)(nil)).Elem(),
		"Hash32": reflect.TypeOf((*hash.Hash32)(nil)).Elem(),
		"Hash64": reflect.TypeOf((*hash.Hash64)(nil)).Elem(),
	}
}
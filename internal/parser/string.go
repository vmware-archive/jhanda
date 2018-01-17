package parser

import (
	"flag"
	"reflect"
)

type String struct{}

func NewString(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (String, error) {
	var defaultValue string
	defaultStr, ok := tags.Lookup("default")
	if ok {
		defaultValue = defaultStr
	}

	short, ok := tags.Lookup("short")
	if ok {
		set.StringVar(field.Addr().Interface().(*string), short, defaultValue, "")
	}

	long, ok := tags.Lookup("long")
	if ok {
		set.StringVar(field.Addr().Interface().(*string), long, defaultValue, "")
	}

	return String{}, nil
}

package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

type Bool struct{}

func NewBool(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (Bool, error) {
	var defaultValue bool
	defaultStr, ok := tags.Lookup("default")
	if ok {
		var err error
		defaultValue, err = strconv.ParseBool(defaultStr)
		if err != nil {
			return Bool{}, fmt.Errorf("could not parse bool default value %q: %s", defaultStr, err)
		}
	}

	short, ok := tags.Lookup("short")
	if ok {
		set.BoolVar(field.Addr().Interface().(*bool), short, defaultValue, "")
	}

	long, ok := tags.Lookup("long")
	if ok {
		set.BoolVar(field.Addr().Interface().(*bool), long, defaultValue, "")
	}

	return Bool{}, nil
}

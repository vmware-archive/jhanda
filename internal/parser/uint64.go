package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

type Uint64 struct{}

func NewUint64(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (Uint64, error) {
	var defaultValue uint64
	defaultStr, ok := tags.Lookup("default")
	if ok {
		var err error
		defaultValue, err = strconv.ParseUint(defaultStr, 0, 64)
		if err != nil {
			return Uint64{}, fmt.Errorf("could not parse uint64 default value %q: %s", defaultStr, err)
		}
	}

	short, ok := tags.Lookup("short")
	if ok {
		set.Uint64Var(field.Addr().Interface().(*uint64), short, defaultValue, "")
	}

	long, ok := tags.Lookup("long")
	if ok {
		set.Uint64Var(field.Addr().Interface().(*uint64), long, defaultValue, "")
	}

	return Uint64{}, nil
}

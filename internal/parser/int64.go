package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

type Int64 struct{}

func NewInt64(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (Int64, error) {
	var defaultValue int64
	defaultStr, ok := tags.Lookup("default")
	if ok {
		var err error
		defaultValue, err = strconv.ParseInt(defaultStr, 0, 64)
		if err != nil {
			return Int64{}, fmt.Errorf("could not parse int64 default value %q: %s", defaultStr, err)
		}
	}

	short, ok := tags.Lookup("short")
	if ok {
		set.Int64Var(field.Addr().Interface().(*int64), short, defaultValue, "")
	}

	long, ok := tags.Lookup("long")
	if ok {
		set.Int64Var(field.Addr().Interface().(*int64), long, defaultValue, "")
	}

	return Int64{}, nil
}

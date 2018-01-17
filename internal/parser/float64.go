package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

type Float64 struct{}

func NewFloat64(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (Float64, error) {
	var defaultValue float64
	defaultStr, ok := tags.Lookup("default")
	if ok {
		var err error
		defaultValue, err = strconv.ParseFloat(defaultStr, 64)
		if err != nil {
			return Float64{}, fmt.Errorf("could not parse float64 default value %q: %s", defaultStr, err)
		}
	}

	short, ok := tags.Lookup("short")
	if ok {
		set.Float64Var(field.Addr().Interface().(*float64), short, defaultValue, "")
	}

	long, ok := tags.Lookup("long")
	if ok {
		set.Float64Var(field.Addr().Interface().(*float64), long, defaultValue, "")
	}

	return Float64{}, nil
}

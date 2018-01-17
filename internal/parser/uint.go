package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

type Uint struct{}

func NewUint(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (Uint, error) {
	var defaultValue uint64
	defaultStr, ok := tags.Lookup("default")
	if ok {
		var err error
		defaultValue, err = strconv.ParseUint(defaultStr, 0, 0)
		if err != nil {
			return Uint{}, fmt.Errorf("could not parse uint default value %q: %s", defaultStr, err)
		}
	}

	short, ok := tags.Lookup("short")
	if ok {
		set.UintVar(field.Addr().Interface().(*uint), short, uint(defaultValue), "")
	}

	long, ok := tags.Lookup("long")
	if ok {
		set.UintVar(field.Addr().Interface().(*uint), long, uint(defaultValue), "")
	}

	return Uint{}, nil
}

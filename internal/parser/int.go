package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

type Int struct{}

func NewInt(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (Int, error) {
	var defaultValue int64
	defaultStr, ok := tags.Lookup("default")
	if ok {
		var err error
		defaultValue, err = strconv.ParseInt(defaultStr, 0, 0)
		if err != nil {
			return Int{}, fmt.Errorf("could not parse int default value %q: %s", defaultStr, err)
		}
	}

	short, ok := tags.Lookup("short")
	if ok {
		set.IntVar(field.Addr().Interface().(*int), short, int(defaultValue), "")
	}

	long, ok := tags.Lookup("long")
	if ok {
		set.IntVar(field.Addr().Interface().(*int), long, int(defaultValue), "")
	}

	return Int{}, nil
}

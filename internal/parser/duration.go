package parser

import (
	"flag"
	"fmt"
	"reflect"
	"time"
)

type Duration struct{}

func NewDuration(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (Duration, error) {
	var defaultValue time.Duration
	defaultStr, ok := tags.Lookup("default")
	if ok {
		var err error
		defaultValue, err = time.ParseDuration(defaultStr)
		if err != nil {
			return Duration{}, fmt.Errorf("could not parse duration default value %q: %s", defaultStr, err)
		}
	}

	short, ok := tags.Lookup("short")
	if ok {
		set.DurationVar(field.Addr().Interface().(*time.Duration), short, defaultValue, "")
	}

	long, ok := tags.Lookup("long")
	if ok {
		set.DurationVar(field.Addr().Interface().(*time.Duration), long, defaultValue, "")
	}

	return Duration{}, nil
}

package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

func NewBool(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (*Flag, error) {
	var defaultValue bool
	parsedTags := newParseTags(tags, set)

	err := parsedTags.setDefault(func(defaultStr string) error {
		var err error
		defaultValue, err = strconv.ParseBool(defaultStr)
		if err != nil {
			return fmt.Errorf("could not parse bool default value %q: %s", defaultStr, err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	parsedTags.setShort(func(short string) {
		set.BoolVar(field.Addr().Interface().(*bool), short, defaultValue, "")
	})

	parsedTags.setAlias(func(alias string) {
		set.BoolVar(field.Addr().Interface().(*bool), alias, defaultValue, "")
	})

	parsedTags.setLong(func(long string) {
		set.BoolVar(field.Addr().Interface().(*bool), long, defaultValue, "")
	})

	err = parsedTags.setEnv(func(envOpt, envStr string) error {
		envValue, err := strconv.ParseBool(envStr)
		if err != nil {
			return fmt.Errorf("could not parse bool environment variable %s value %q: %s", envOpt, envStr, err)
		}

		field.SetBool(envValue)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return parsedTags.getFlag(), nil
}

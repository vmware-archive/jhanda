package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

func NewFloat64(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (*Flag, error) {
	var defaultValue float64
	parsedTags := newParseTags(tags, set)

	err := parsedTags.setDefault(func(defaultStr string) error {
		var err error
		defaultValue, err = strconv.ParseFloat(defaultStr, 64)
		if err != nil {
			return fmt.Errorf("could not parse float64 default value %q: %s", defaultStr, err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	parsedTags.setShort(func(short string) {
		set.Float64Var(field.Addr().Interface().(*float64), short, defaultValue, "")
	})

	parsedTags.setAlias(func(alias string) {
		set.Float64Var(field.Addr().Interface().(*float64), alias, defaultValue, "")
	})

	parsedTags.setLong(func(long string) {
		set.Float64Var(field.Addr().Interface().(*float64), long, defaultValue, "")
	})

	err = parsedTags.setEnv(func(envOpt, envStr string) error {
		envValue, err := strconv.ParseFloat(envStr, 64)
		if err != nil {
			return fmt.Errorf("could not parse float64 environment variable %s value %q: %s", envOpt, envStr, err)
		}

		field.SetFloat(envValue)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return parsedTags.getFlag(), nil
}

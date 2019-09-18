package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

func NewUint(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (*Flag, error) {
	var defaultValue uint64
	parsedTags := newParseTags(tags, set)

	err := parsedTags.setDefault(func(defaultStr string) error {
		var err error
		defaultValue, err = strconv.ParseUint(defaultStr, 0, 0)
		if err != nil {
			return fmt.Errorf("could not parse uint default value %q: %s", defaultStr, err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	parsedTags.setShort(func(short string) {
		set.UintVar(field.Addr().Interface().(*uint), short, uint(defaultValue), "")
	})

	parsedTags.setAlias(func(alias string) {
		set.UintVar(field.Addr().Interface().(*uint), alias, uint(defaultValue), "")
	})

	parsedTags.setLong(func(long string) {
		set.UintVar(field.Addr().Interface().(*uint), long, uint(defaultValue), "")
	})

	err = parsedTags.setEnv(func(envOpt, envStr string) error {
		envValue, err := strconv.ParseUint(envStr, 0, 0)
		if err != nil {
			return fmt.Errorf("could not parse uint environment variable %s value %q: %s", envOpt, envStr, err)
		}

		field.SetUint(envValue)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return parsedTags.getFlag(), nil
}

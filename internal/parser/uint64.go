package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

func NewUint64(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (*Flag, error) {
	var defaultValue uint64
	parsedTags := newParseTags(tags, set)

	err := parsedTags.setDefault(func(defaultStr string) error {
		var err error
		defaultValue, err = strconv.ParseUint(defaultStr, 0, 64)
		if err != nil {
			return fmt.Errorf("could not parse uint64 default value %q: %s", defaultStr, err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	parsedTags.setShort(func(short string) {
		set.Uint64Var(field.Addr().Interface().(*uint64), short, defaultValue, "")
	})

	parsedTags.setAlias(func(alias string) {
		set.Uint64Var(field.Addr().Interface().(*uint64), alias, defaultValue, "")
	})

	parsedTags.setLong(func(long string) {
		set.Uint64Var(field.Addr().Interface().(*uint64), long, defaultValue, "")
	})

	err = parsedTags.setEnv(func(envOpt, envStr string) error {
		envValue, err := strconv.ParseUint(envStr, 0, 64)
		if err != nil {
			return fmt.Errorf("could not parse uint64 environment variable %s value %q: %s", envOpt, envStr, err)
		}

		field.SetUint(envValue)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return parsedTags.getFlag(), nil
}

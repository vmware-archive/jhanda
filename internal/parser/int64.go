package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

func NewInt64(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (*Flag, error) {
	var defaultValue int64
	parsedTags := newParseTags(tags, set)

	err := parsedTags.setDefault(func(defaultStr string) error {
		var err error
		defaultValue, err = strconv.ParseInt(defaultStr, 0, 64)
		if err != nil {
			return fmt.Errorf("could not parse int64 default value %q: %s", defaultStr, err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	parsedTags.setShort(func(short string) {
		set.Int64Var(field.Addr().Interface().(*int64), short, defaultValue, "")
	})

	parsedTags.setAlias(func(alias string) {
		set.Int64Var(field.Addr().Interface().(*int64), alias, defaultValue, "")
	})

	parsedTags.setLong(func(long string) {
		set.Int64Var(field.Addr().Interface().(*int64), long, defaultValue, "")
	})

	err = parsedTags.setEnv(func(envOpt, envStr string) error {
		envValue, err := strconv.ParseInt(envStr, 0, 64)
		if err != nil {
			return fmt.Errorf("could not parse int64 environment variable %s value %q: %s", envOpt, envStr, err)
		}

		field.SetInt(envValue)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return parsedTags.getFlag(), nil
}

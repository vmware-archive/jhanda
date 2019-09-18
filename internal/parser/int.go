package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

func NewInt(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (*Flag, error) {
	var defaultValue int64
	parsedTags := newParseTags(tags, set)

	err := parsedTags.setDefault(func(defaultStr string) error {
		var err error
		defaultValue, err = strconv.ParseInt(defaultStr, 0, 0)
		if err != nil {
			return fmt.Errorf("could not parse int default value %q: %s", defaultStr, err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	parsedTags.setShort(func(short string) {
		set.IntVar(field.Addr().Interface().(*int), short, int(defaultValue), "")
	})

	parsedTags.setAlias(func(alias string) {
		set.IntVar(field.Addr().Interface().(*int), alias, int(defaultValue), "")
	})

	parsedTags.setLong(func(long string) {
		set.IntVar(field.Addr().Interface().(*int), long, int(defaultValue), "")
	})

	err = parsedTags.setEnv(func(envOpt, envStr string) error {
		envValue, err := strconv.ParseInt(envStr, 0, 0)
		if err != nil {
			return fmt.Errorf("could not parse int environment variable %s value %q: %s", envOpt, envStr, err)
		}

		field.SetInt(envValue)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return parsedTags.getFlag(), nil
}

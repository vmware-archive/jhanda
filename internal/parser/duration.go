package parser

import (
	"flag"
	"fmt"
	"reflect"
	"time"
)

func NewDuration(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (*Flag, error) {
	var defaultValue time.Duration
	parsedTags := newParseTags(tags, set)

	err := parsedTags.setDefault(func(defaultStr string) error {
		var err error
		defaultValue, err = time.ParseDuration(defaultStr)
		if err != nil {
			return fmt.Errorf("could not parse duration default value %q: %s", defaultStr, err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	parsedTags.setShort(func(short string) {
		set.DurationVar(field.Addr().Interface().(*time.Duration), short, defaultValue, "")
	})

	parsedTags.setAlias(func(alias string) {
		set.DurationVar(field.Addr().Interface().(*time.Duration), alias, defaultValue, "")
	})

	parsedTags.setLong(func(long string) {
		set.DurationVar(field.Addr().Interface().(*time.Duration), long, defaultValue, "")
	})

	err = parsedTags.setEnv(func(envOpt, envStr string) error {
		envValue, err := time.ParseDuration(envStr)
		if err != nil {
			return fmt.Errorf("could not parse duration environment variable %s value %q: %s", envOpt, envStr, err)
		}

		field.SetInt(int64(envValue))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return parsedTags.getFlag(), nil
}

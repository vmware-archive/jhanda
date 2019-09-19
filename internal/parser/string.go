package parser

import (
	"flag"
	"os"
	"reflect"
	"strings"
)

func NewString(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (*Flag, error) {
	var defaultValue string
	parsedTags := newParseTags(tags, set)

	err := parsedTags.setDefault(func(defaultStr string) error {
		defaultValue = defaultStr
		return nil
	})

	if err != nil {
		return nil, err
	}

	parsedTags.setShort(func(short string) {
		set.StringVar(field.Addr().Interface().(*string), short, defaultValue, "")
	})

	parsedTags.setAlias(func(alias string) {
		set.StringVar(field.Addr().Interface().(*string), alias, defaultValue, "")
	})

	parsedTags.setLong(func(long string) {
		set.StringVar(field.Addr().Interface().(*string), long, defaultValue, "")
	})

	f := parsedTags.getFlag()
	env, ok := tags.Lookup("env")
	if ok {
		envOpts := strings.Split(env, ",")

		for _, envOpt := range envOpts {
			envStr, ok := os.LookupEnv(envOpt)
			if ok {
				field.SetString(envStr)
				f.set = true
				break
			}
		}
	}

	return f, nil
}

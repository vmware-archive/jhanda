package parser

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func NewSlice(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (*Flag, error) {
	collection := field.Addr().Interface().(*[]string)

	parsedTags := newParseTags(tags, set)

	_ = parsedTags.setDefault(func(defaultStr string) error {
		separated := strings.Split(defaultStr, ",")
		*collection = append(*collection, separated...)
		return nil
	})

	slice := StringSlice{collection}

	parsedTags.setShort(func(short string) {
		set.Var(&slice, short, "")
	})

	parsedTags.setAlias(func(alias string) {
		set.Var(&slice, alias, "")
	})

	parsedTags.setLong(func(long string) {
		set.Var(&slice, long, "")
	})

	f := parsedTags.getFlag()
	env, ok := tags.Lookup("env")
	if ok {
		envOpts := strings.Split(env, ",")

		for _, envOpt := range envOpts {
			envStr, ok := os.LookupEnv(envOpt)
			if ok {
				separated := strings.Split(envStr, ",")
				*collection = append(*collection, separated...)
				f.set = true
				break
			}
		}
	}

	return f, nil
}

type StringSlice struct {
	slice *[]string
}

func (ss *StringSlice) String() string {
	return fmt.Sprintf("%s", ss.slice)
}

func (ss *StringSlice) Set(item string) error {
	*ss.slice = append(*ss.slice, item)
	return nil
}

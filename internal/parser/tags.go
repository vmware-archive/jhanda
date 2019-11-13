package parser

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type parsedTags struct {
	tags reflect.StructTag
	flag *Flag
	set *flag.FlagSet
}

func newParseTags(tags reflect.StructTag, set *flag.FlagSet) parsedTags {
	return parsedTags{
		tags: tags,
		flag: &Flag{},
		set: set,
	}
}

func (t parsedTags) getDefault() (string, bool) {
	return t.tags.Lookup("default")
}

func (t parsedTags) setDefault(fn func(string) error) error {
	value, exists := t.tags.Lookup("default")
	if exists {
		return fn(value)
	}
	return nil
}

func (t parsedTags) setShort(fn func(string)) {
	short, exists := t.tags.Lookup("short")
	if exists {
		fn(short)
		t.flag.flags = append(t.flag.flags, t.set.Lookup(short))
		t.flag.name = fmt.Sprintf("-%s", short)
	}
}

func (t parsedTags) setLong(fn func(string)) {
	long, exists := t.tags.Lookup("long")
	if exists {
		fn(long)
		t.flag.flags = append(t.flag.flags, t.set.Lookup(long))
		t.flag.name = fmt.Sprintf("--%s", long)
	}
}

func (t parsedTags) setAlias(fn func(string)) {
	alias, exists := t.tags.Lookup("alias")
	if exists {
		aliases := strings.Split(alias, ",")
		for _, a := range aliases {
			fn(a)
			t.flag.flags = append(t.flag.flags, t.set.Lookup(a))
			t.flag.name = fmt.Sprintf("--%s", a)
		}
	}
}

func (t parsedTags) setEnv(fn func(string, string) error) error {
	env, exists := t.tags.Lookup("env")
	if exists {
		envOpts := strings.Split(env, ",")

		for _, envOpt := range envOpts {
			envStr := os.Getenv(envOpt)
			if envStr != "" {
				err := fn(envOpt, envStr)
				if err != nil {
					return err
				}
				t.flag.set = true
				return nil
			}
		}
	}
	return nil
}

func (t parsedTags) getFlag() *Flag {
	_, t.flag.required = t.tags.Lookup("required")

	return t.flag
}

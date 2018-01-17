package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strings"
)

type Slice struct{}

func NewSlice(set *flag.FlagSet, field reflect.Value, tags reflect.StructTag) (Slice, error) {
	collection := field.Addr().Interface().(*[]string)

	defaultSlice, ok := tags.Lookup("default")
	if ok {
		separated := strings.Split(defaultSlice, ",")
		*collection = append(*collection, separated...)
	}

	slice := StringSlice{collection}

	short, ok := tags.Lookup("short")
	if ok {
		set.Var(&slice, short, "")
	}

	long, ok := tags.Lookup("long")
	if ok {
		set.Var(&slice, long, "")
	}

	return Slice{}, nil
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

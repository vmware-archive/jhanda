package parser

import (
	"flag"
	"fmt"
	"reflect"
	"strings"
)

type Slice struct {
	Set   *flag.FlagSet
	Field reflect.Value
	Tags  reflect.StructTag
}

func (s Slice) Execute() error {
	collection := s.Field.Addr().Interface().(*StringSlice)

	defaultSlice, ok := s.Tags.Lookup("default")
	if ok {
		separated := strings.Split(defaultSlice, ",")
		*collection = append(*collection, separated...)
	}

	short, ok := s.Tags.Lookup("short")
	if ok {
		s.Set.Var(collection, short, "")
	}

	long, ok := s.Tags.Lookup("long")
	if ok {
		s.Set.Var(collection, long, "")
	}

	return nil
}

type StringSlice []string

func (ss *StringSlice) String() string {
	return fmt.Sprintf("%s", *ss)
}

func (ss *StringSlice) Set(item string) error {
	*ss = append(*ss, item)
	return nil
}

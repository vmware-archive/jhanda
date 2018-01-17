package jhanda

import (
	"flag"
	"fmt"
	"io/ioutil"
	"reflect"
	"time"

	"github.com/pivotal-cf/jhanda/internal/parser"
)

func Parse(receiver interface{}, args []string) ([]string, error) {
	set := flag.NewFlagSet("", flag.ContinueOnError)
	set.SetOutput(ioutil.Discard)
	set.Usage = func() {}

	ptr := reflect.ValueOf(receiver)
	if ptr.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("unexpected non-pointer type %s for flag receiver", ptr.Kind())
	}

	v := ptr.Elem()
	t := v.Type()

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("unexpected pointer to non-struct type %s", t.Kind())
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		var err error

		switch {
		case field.Type.Kind() == reflect.Bool:
			_, err = parser.NewBool(set, v.Field(i), field.Tag)

		case field.Type.Kind() == reflect.Float64:
			_, err = parser.NewFloat64(set, v.Field(i), field.Tag)

		case field.Type == reflect.TypeOf(time.Duration(0)):
			_, err = parser.NewDuration(set, v.Field(i), field.Tag)

		case field.Type.Kind() == reflect.Int64:
			_, err = parser.NewInt64(set, v.Field(i), field.Tag)

		case field.Type.Kind() == reflect.Int:
			_, err = parser.NewInt(set, v.Field(i), field.Tag)

		case field.Type.Kind() == reflect.String:
			_, err = parser.NewString(set, v.Field(i), field.Tag)

		case field.Type.Kind() == reflect.Uint64:
			_, err = parser.NewUint64(set, v.Field(i), field.Tag)

		case field.Type.Kind() == reflect.Uint:
			_, err = parser.NewUint(set, v.Field(i), field.Tag)

		case field.Type.Kind() == reflect.Slice:
			_, err = parser.NewSlice(set, v.Field(i), field.Tag)

		default:
			return nil, fmt.Errorf("unexpected flag receiver field type %s", field.Type.Kind())
		}
		if err != nil {
			return nil, err
		}
	}

	err := set.Parse(args)
	if err != nil {
		return nil, err
	}

	return set.Args(), nil
}

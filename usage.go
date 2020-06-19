package jhanda

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Usage provides all of the details describing a Command, including a
// description, a shorter description (used when display a list of commands),
// and the flag options offered by the Command.
type Usage struct {
	Description      string
	ShortDescription string
	Flags            interface{}
}

// PrintUsage will return a string representation of the options provided by a
// Command flag set.
func PrintUsage(receiver interface{}) (string, error) {
	v := reflect.ValueOf(receiver)
	t := v.Type()
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("unexpected pointer to non-struct type %s", t.Kind())
	}

	fields := getFields(t)

	var usage []string
	var length int
	for _, field := range fields {
		var arguments []string
		long, ok := field.Tag.Lookup("long")
		if ok {
			arguments = append(arguments, fmt.Sprintf("--%s", long))
		}

		short, ok := field.Tag.Lookup("short")
		if ok {
			arguments = append(arguments, fmt.Sprintf("-%s", short))
		}

		envs, ok := field.Tag.Lookup("env")
		if ok {
			for _, env := range strings.Split(envs, ",") {
				arguments = append(arguments, fmt.Sprintf("%s", env))
			}
		}

		field := strings.Join(arguments, ", ")

		if len(field) > length {
			length = len(field)
		}

		usage = append(usage, field)
	}

	for i, line := range usage {
		usage[i] = pad(line, " ", length)
	}

	for i, field := range fields {
		var kindParts []string
		if _, ok := field.Tag.Lookup("required"); ok {
			kindParts = append(kindParts, "required")
		}

		kind := field.Type.Kind().String()
		if kind == reflect.Slice.String() {
			kind = field.Type.Elem().Kind().String()
			kindParts = append(kindParts, "variadic")
		}

		if len(kindParts) > 0 {
			kind = fmt.Sprintf("%s (%s)", kind, strings.Join(kindParts, ", "))
		}

		line := fmt.Sprintf("%s  %s", usage[i], kind)

		if len(line) > length {
			length = len(line)
		}

		usage[i] = line
	}

	for i, line := range usage {
		usage[i] = pad(line, " ", length)
	}

	for i, field := range fields {
		description, ok := field.Tag.Lookup("description")
		if ok {
			if _, ok := field.Tag.Lookup("deprecated"); ok {
				description = fmt.Sprintf("**DEPRECATED** %s", description)
			}

			if _, ok := field.Tag.Lookup("experimental"); ok {
				description = fmt.Sprintf("**EXPERIMENTAL** %s", description)
			}

			usage[i] += fmt.Sprintf("  %s", description)
		}
	}

	for i, field := range fields {
		defaultValue, ok := field.Tag.Lookup("default")
		if ok {
			usage[i] += fmt.Sprintf(" (default: %s)", defaultValue)
		}
	}

	for i, field := range fields {
		aliases, ok := field.Tag.Lookup("alias")
		if ok {
			var arguments []string
			for _, alias := range strings.Split(aliases, ",") {
				arguments = append(arguments, fmt.Sprintf("--%s", alias))
			}
			usage[i] += fmt.Sprintf("\n  (aliases: %s)", strings.Join(arguments, ", "))
		}
	}

	sort.Strings(usage)

	return strings.Join(usage, "\n"), nil
}

func getFields(t reflect.Type) []reflect.StructField {
	var fields []reflect.StructField
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct {
			fields = append(fields, getFields(field.Type)...)
			continue
		}

		fields = append(fields, field)
	}
	return fields
}

func pad(str, pad string, length int) string {
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}

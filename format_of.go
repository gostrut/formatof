package formatof

import (
	"fmt"
	"reflect"
	"regexp"

	"gopkg.in/gostrut/strut.v1/invalid"
)

// Validator for formatof validates a string field against a regexp. Empty
// strings will be skipped. This does not check presence of.
// Regex must compile or an error is returned
func Validator(n, t string, v *reflect.Value) (invalid.Field, error) {
	s, ok := v.Interface().(string)
	if !ok {
		return nil,
			fmt.Errorf("format_of cannot be applied to %s, it is not a string", n)
	}
	if "" == s {
		return nil, nil
	}

	r, err := regexp.Compile(t)
	if err != nil {
		return nil, err
	}

	if !r.MatchString(s) {
		f := field{
			name:   n,
			regStr: r.String(),
		}

		return f, nil
	}

	return nil, nil
}

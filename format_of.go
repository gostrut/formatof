package formatof

import "fmt"
import "reflect"
import "regexp"
import "github.com/gostrut/invalid"

// Validator for formatof validates a string field against a regexp. Empty
// strings will be skipped. This does not check presence of.
// Regex must compile or an error is returned
func Validator(name, tagStr string, vo *reflect.Value) (invalid.Field, error) {
	if vo.Kind() != reflect.String {
		return nil, fmt.Errorf("format_of cannot be applied to %s, it is not a string", name)
	}

	str := vo.Interface().(string)
	if "" == str {
		return nil, nil
	}

	reg, err := regexp.Compile(tagStr)
	if err != nil {
		return nil, err
	}

	if !reg.MatchString(str) {
		f := iField{
			name:   name,
			regStr: reg.String(),
		}
		return f, nil
	}

	return nil, nil
}

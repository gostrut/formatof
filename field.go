package formatof

import (
	"fmt"
)

type field struct {
	name   string
	regStr string
}

func (f field) Name() string {
	return f.name
}

func (f field) Validator() string {
	return "FormatOf"
}

func (f field) Error() string {
	return fmt.Sprintf("%s does not match `%s`", f.Name(), f.regStr)
}

package formatof

import "fmt"

type iField struct {
	name   string
	regStr string
}

func (f iField) Name() string {
	return f.name
}

func (f iField) Validator() string {
	return "FormatOf"
}

func (f iField) Error() string {
	return fmt.Sprintf("%s does not match `%s`", f.Name(), f.regStr)
}

package formatof

import "testing"
import "github.com/nowk/assert"
import "github.com/gostrut/strut"

func TestFormatOf(t *testing.T) {
	val := strut.NewValidator()
	val.Checks("format_of", Validator)

	type Person struct {
		Email string `format_of:"^\\w+@\\w+\\.[a-z]{2,4}$"`
	}

	a := Person{Email: "notanemail"}
	b := Person{Email: "almost@there.comtoolong"}

	for _, v := range []struct {
		i interface{}
	}{
		{a},
		{b},
	} {
		fields, err := val.Validates(v.i)
		f := fields.Get("Email")[0]
		assert.Nil(t, err)
		assert.False(t, fields.Valid())
		assert.Equal(t, "Email does not match `^\\w+@\\w+\\.[a-z]{2,4}$`", f.Error())
	}

	e := Person{
		Email: "email@company.com",
	}

	fields, err := val.Validates(e)
	assert.Nil(t, err)
	assert.True(t, fields.Valid())
}

func TestFormatOfError(t *testing.T) {
	val := strut.NewValidator()
	val.Checks("format_of", Validator)

	type Person struct {
		Email string `format_of:"a(b"`
	}

	type User struct {
		Person Person `format_of:"^person$"`
	}

	a := User{Person: Person{}}
	b := Person{Email: "email@company.com"}

	for _, v := range []struct {
		i interface{}
		e string
	}{
		{a, "format_of cannot be applied to Person, it is not a string"},
		{b, "error parsing regexp: missing closing ): `a(b`"},
	} {
		_, err := val.Validates(v.i)
		assert.Equal(t, err.Error(), v.e)
	}
}

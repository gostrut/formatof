# formatof

[![Build Status](https://travis-ci.org/gostrut/formatof.svg?branch=master)](https://travis-ci.org/gostrut/formatof)

Validate format of

## Install

    go get gopkg.in/gostrut/formatof.v1

## Example

    type Person struct {
      Email string `format_of:"^\\w+@\\w+\\.[a-z]{2,4}$`"`
    }

    val := NewValidator()
    val.Add("format_of", formatof.Validator)

    p := Person{Email: "badnotanemail"}
    fields, err := val.Check(p)
    if err != nil {
      // handle error
    }
    if !fields.Valid() {
      // handle invalid fields
    }

## License

MIT

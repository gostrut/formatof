# formatof

Validate format of

## Example

    type Person struct {
      Email string `format_of:"^\\w+@\\w+\\.[a-z]{2,4}$`"`
    }

    val := NewValidator()
    val.Checks("format_of", formatof.Validator)

    p := Person{Email: "badnotanemail"}
    fields, err := val.Validates(p)

## License

MIT
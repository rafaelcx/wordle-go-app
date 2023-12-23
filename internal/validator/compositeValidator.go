package validator

func Validate(input string) error {
	var validators []Validator

	validators = append(validators, inputSizeValidator{input: input})
	validators = append(validators, inputExistenceValidator{input: input})

	for _, validator := range validators {
		err := validator.validate()

		if err != nil {
			return err
		}
	}
	return nil
}

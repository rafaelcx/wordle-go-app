package validator

func Validate(input string) error {
	for _, validator := range getValidatorList(input) {
		err := validator.validate()

		if err != nil {
			return err
		}
	}
	return nil
}

func getValidatorList(input string) []Validator {
	var validators []Validator
	validators = append(validators, inputSizeValidator{input: input})
	validators = append(validators, inputExistenceValidator{input: input})

	return validators
}

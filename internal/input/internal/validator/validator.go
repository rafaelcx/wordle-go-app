package validator

import "fmt"

type validator interface {
	validate(s string) error
}

func Validate(s string) error {
	for _, validatorList := range getValidatorList() {
		err := validatorList.validate(s)
		if err != nil {
			fmt.Println(err.Error() + "\n")
			return err
		}
	}
	return nil
}

func getValidatorList() []validator {
	return []validator{
		wordLengthValidator{},
	}
}

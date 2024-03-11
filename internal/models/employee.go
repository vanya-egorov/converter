package models

import (
	"errors"
	"regexp"
)

type Employee struct {
	Id       int
	Name     string
	Phone    string
	Birthday *string
}

func (e *Employee) Validate() error {
	if e.Name == "" {
		return errors.New("Имя не может быть пустым")
	}

	phonePattern := `^\+[0-9]{11}$`
	matched, _ := regexp.MatchString(phonePattern, e.Phone)
	if !matched {
		return errors.New("Некорректный формат телефонного номера")
	}

	return nil
}

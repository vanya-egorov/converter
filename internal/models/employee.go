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

// Validate проводит проверку валидности полей Employee
func (e *Employee) Validate() error {
    // Проверка наличия имени
    if e.Name == "" {
        return errors.New("Имя не может быть пустым")
    }

    // Проверка телефонного номера с помощью регулярного выражения
    phonePattern := ^\+[0-9]{11}$
    matched, _ := regexp.MatchString(phonePattern, e.Phone)
    if !matched {
        return errors.New("Некорректный формат телефонного номера")
    }

    return nil
}
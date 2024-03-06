package models

type Employee struct {
	Id       int
	Name     string  // Василий
	Phone    string  // +71234567890
	Birthday *string // RFC3339
}

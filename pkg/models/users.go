package models

type Users struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName,omitempty"`
}

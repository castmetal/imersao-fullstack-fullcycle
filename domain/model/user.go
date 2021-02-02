package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

type User struct {
	Base
	ID string			`json:"id" valid:"uuid"`
	Name string 		`json:"name" valid:"notnull"`
	Email string 		`json:"email" valid:"notnull"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return errors.New("ID, Name and Email are required fields")
	}
	return nil
}
package models

import (
	u "gopher-rest/pkg/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

type Organization struct {
	Id         int `gorm:"primaryKey"`
	Name       string
	Created_By int
	Created_At time.Time
	Updated_At time.Time
	Users      []User `gorm:"many2many:org_users;"`
}

func (org *Organization) Create() map[string]interface{} {

	validate := validator.New()

	if err := validate.Struct(org); err != nil {
		return u.Message(false, "Struct Validation Error.")
	}

	GetDB().Create(org)

	response := u.Message(true, "Organization created successfully.")

	response["organization"] = org

	return response
}

// func (org *Organization) AddUser() map[string]interface{} {

// }
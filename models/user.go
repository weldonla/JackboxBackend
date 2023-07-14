package models

import (
	"github.com/google/uuid"
	"github.com/weldonla/FourLeafPortalApi/datatypes"
	"gorm.io/gorm"
)

type User struct {
	Id        datatypes.MYTYPE `gorm:"primary_key;"`
	FirstName string           `json:"firstName" db:"first_name"`
	LastName  string           `json:"lastName" db:"last_name"`
	UserName  string           `json:"userName" db:"user_name" gorm:"unique"`
	Email     string           `json:"email" db:"email" gorm:"unique"`
	Phone     string           `json:"phone" db:"phone"`
	Password  []byte           `json:"-"`
	IsAdmin   bool             `json:"isAdmin" db:"is_admin"`
}

// BeforeCreate ->
func (p *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	p.Id = datatypes.MYTYPE(id)
	return err
}

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Role  Role
	Name  string
	Email string
}

type Role int

const (
	Customer Role = 0
	Support  Role = 1
	Admin    Role = 2
	Vendor   Role = 3
)

//Package models has the database models used in the application.
package models

import "gorm.io/gorm"

//User model
type User struct {
	gorm.Model
	//Role of the user
	Role Role
	//Name of the user
	Name string
	//Email of the user
	Email    string
	Password string
}

//Role of the user
type Role int

const (
	//Customer denotes the user is a customer
	Customer Role = 0
	//Support denotes the user is a support
	Support Role = 1
	//Admin denotes the user is an admin
	Admin Role = 2
	//Vendor denotes the user is a vendor
	Vendor Role = 3
)

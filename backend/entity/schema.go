package entity

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string `valid:"required~Name cannot be null"`
	Email      string
	EmployeeID string
}

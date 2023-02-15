package entity

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string `valid:"required~Name cannot be null"`
	Email      string
	EmployeeID string `valid:"matches(^([J|M|S])([0-9]{8})$)~EmployeeID must be JMS_12345678"`
}

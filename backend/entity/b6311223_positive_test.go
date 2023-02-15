package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestEmployeePass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	e := Employee{
		Name:       "employee",
		Email:      "employee@gmail.com",
		EmployeeID: "J63112232",
	}

	ok, err := govalidator.ValidateStruct(e)
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())

}

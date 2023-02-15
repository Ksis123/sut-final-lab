package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestEmployeeID(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("EmployeeID must be JMS_12345678", func(t *testing.T) {
		e := Employee{
			Name:       "employee",
			Email:      "employee@gmail.com",
			EmployeeID: "AAA112232",
		}

		ok, err := govalidator.ValidateStruct(e)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("EmployeeID must be JMS_12345678"))
	})
}

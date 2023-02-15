package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestEmployeeNotNull(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Name cannot be null", func(t *testing.T) {
		e := Employee{
			Name:       "",
			Email:      "employee@gmail.com",
			EmployeeID: "J63112232",
		}

		ok, err := govalidator.ValidateStruct(e)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Name cannot be null"))
	})
}

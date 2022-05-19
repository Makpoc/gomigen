package model_test

import (
	"github.com/Makpoc/gomigen/internal/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Middleware", func() {
	Describe("NewMiddleware", func() {
		It("adds the required imports on creation", func() {
			middleware := model.NewMiddleware("Interface")
			impStatements := middleware.Imports.ImportStatements()

			Expect(impStatements).To(HaveLen(2))

			Expect(impStatements).To(ContainElement(HaveField("Alias", "")))
			Expect(impStatements).To(ContainElement(HaveField("Package", "context")))

			Expect(impStatements).To(ContainElement(HaveField("Alias", "")))
			Expect(impStatements).To(ContainElement(HaveField("Package", "github.com/Makpoc/gomigen/types")))
		})
	})
})

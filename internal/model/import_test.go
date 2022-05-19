package model_test

import (
	"github.com/Makpoc/gomigen/internal/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Imports", func() {
	DescribeTable("Add",
		func(givenExistingImports []model.Import, givenAddImport model.Import, expectedAlias string) {
			imports := model.NewImports()
			for _, imp := range givenExistingImports {
				imports.Add(imp)
			}
			actualAlias := imports.Add(givenAddImport)
			Expect(actualAlias).To(Equal(expectedAlias))
		},
		Entry("no custom alias when new and only import is added",
			[]model.Import{}, model.Import{
				Alias:   "log",
				Package: "log_pkg",
			},
			"log",
		),
		Entry("a custom alias when an import with the given alias exists",
			[]model.Import{
				{
					Alias:   "log",
					Package: "log_pkg",
				},
			},
			model.Import{
				Alias:   "log",
				Package: "log_pkg/v2",
			},
			"log0",
		),
	)

	DescribeTable("ImportStatements",
		func(givenImports []model.Import, expectedImports []model.Import) {
			imports := model.NewImports()
			for _, imp := range givenImports {
				imports.Add(imp)
			}
			impStatements := imports.ImportStatements()
			Expect(impStatements).To(HaveLen(len(givenImports)))

			for _, expectedImports := range expectedImports {
				Expect(impStatements).To(ContainElement(HaveField("Alias", expectedImports.Alias)))
				Expect(impStatements).To(ContainElement(HaveField("Package", expectedImports.Package)))
			}
		},
		Entry("there are no imports", []model.Import{}, []model.Import{}),
		Entry("there are imports without alias conflicts",
			[]model.Import{
				{
					Alias: "log", Package: "log/v1",
				}, {
					Alias: "correlation", Package: "repository.com/correlation",
				},
			}, []model.Import{
				{
					Alias: "", Package: "log/v1",
				}, {
					Alias: "", Package: "repository.com/correlation",
				},
			},
		),
		Entry("there are imports with alias conflicts",
			[]model.Import{
				{
					Alias: "log", Package: "log/v1",
				}, {
					Alias: "log", Package: "log/v2",
				}, {
					Alias: "correlation", Package: "repository.com/correlation",
				},
			}, []model.Import{
				{
					Alias: "", Package: "log/v1",
				}, {
					Alias: "log0", Package: "log/v2",
				}, {
					Alias: "", Package: "repository.com/correlation",
				},
			},
		),
		Entry("there are imports with multiple alias conflicts",
			[]model.Import{
				{
					Alias: "log", Package: "log/v1",
				}, {
					Alias: "log", Package: "log/v3",
				}, {
					Alias: "log", Package: "log/v2",
				},
			}, []model.Import{
				{
					Alias: "", Package: "log/v1",
				}, {
					Alias: "log0", Package: "log/v3",
				}, {
					Alias: "log1", Package: "log/v2",
				},
			},
		),
	)
})

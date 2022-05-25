package model_test

import (
	"github.com/Makpoc/gomigen/internal/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Return", func() {
	DescribeTable("LastVarTypeIsError",
		func(params model.Params, expectedResult bool) {
			method := new(model.Method)
			for _, p := range params {
				method.AddReturn(p)
			}
			Expect(method.Returns.LastVarTypeIsError()).To(Equal(expectedResult))
		},
		Entry("there are no returns", model.Params{}, false),
		Entry("there is one return - not a error", model.Params{
			{
				Name: "",
				Type: "string",
			},
		}, false),
		Entry("there is one return - an error", model.Params{
			{
				Name: "",
				Type: "error",
			},
		}, true),
		Entry("there are two returns - no error", model.Params{
			{
				Name: "",
				Type: "string",
			}, {
				Name: "",
				Type: "int",
			},
		}, false),
		Entry("there are two returns - first is error", model.Params{
			{
				Name: "",
				Type: "error",
			}, {
				Name: "",
				Type: "int",
			},
		}, false),
		Entry("there are two returns - last is error", model.Params{
			{
				Name: "",
				Type: "string",
			}, {
				Name: "",
				Type: "error",
			},
		}, true),
	)

	DescribeTable("ErrorVarName",
		func(params model.Params, expectedOutput string) {
			method := new(model.Method)
			for _, p := range params {
				method.AddReturn(p)
			}
			Expect(method.Returns.ErrorVarName()).To(Equal(expectedOutput))
		},
		Entry("there are no returns", model.Params{}, ""),
		Entry("there is one return - not a error", model.Params{
			{
				Name: "",
				Type: "string",
			},
		}, ""),
		Entry("there is one return - an error", model.Params{
			{
				Name: "",
				Type: "error",
			},
		}, "res0"),
		Entry("there are two returns - no error", model.Params{
			{
				Name: "",
				Type: "string",
			}, {
				Name: "",
				Type: "int",
			},
		}, ""),
		Entry("there are two returns - first is error", model.Params{
			{
				Name: "",
				Type: "error",
			}, {
				Name: "",
				Type: "int",
			},
		}, ""),
		Entry("there are two returns - last is error", model.Params{
			{
				Name: "",
				Type: "string",
			}, {
				Name: "",
				Type: "error",
			},
		}, "res1"),
		Entry("there are two named returns - last is error", model.Params{
			{
				Name: "str1",
				Type: "string",
			}, {
				Name: "err",
				Type: "error",
			},
		}, "res1"),
	)

	DescribeTable("ForMethodSignature",
		func(params model.Params, expectedOutput string) {
			method := new(model.Method)
			for _, p := range params {
				method.AddReturn(p)
			}
			Expect(method.Returns.ForMethodSignature()).To(Equal(expectedOutput))
		},
		Entry("there are no returns", model.Params{}, ""),
		Entry("there is one unnamed return", model.Params{
			{
				Name: "",
				Type: "string",
			},
		}, "string"),
		Entry("there is one named return", model.Params{
			{
				Name: "err",
				Type: "error",
			},
		}, "error /* err */"),
		Entry("there are two unnamed returns", model.Params{
			{
				Name: "",
				Type: "string",
			}, {
				Name: "",
				Type: "error",
			},
		}, "string, error"),
		Entry("there are two named returns", model.Params{
			{
				Name: "str",
				Type: "string",
			}, {
				Name: "err",
				Type: "error",
			},
		}, "string /* str */, error /* err */"),
	)
	DescribeTable("ReturnVarNames",
		func(params model.Params, expectedOutput string) {
			method := new(model.Method)
			for _, p := range params {
				method.AddReturn(p)
			}
			Expect(method.Returns.ReturnVarNames()).To(Equal(expectedOutput))
		},
		Entry("there are no returns", model.Params{}, ""),
		Entry("there is one unnamed return", model.Params{
			{
				Name: "",
				Type: "string",
			},
		}, "res0"),
		Entry("there is one named return", model.Params{
			{
				Name: "str",
				Type: "string",
			},
		}, "res0"),
		Entry("there are two unnamed returns", model.Params{
			{
				Name: "",
				Type: "string",
			}, {
				Name: "",
				Type: "error",
			},
		}, "res0, res1"),
		Entry("there are two unnamed returns", model.Params{
			{
				Name: "str",
				Type: "string",
			}, {
				Name: "err",
				Type: "error",
			},
		}, "res0, res1"),
	)
})

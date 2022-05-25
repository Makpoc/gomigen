package model_test

import (
	"github.com/Makpoc/gomigen/internal/model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Argument", func() {
	DescribeTable("ForMethodSignature",
		func(params model.Params, expectedOutput string) {
			method := new(model.Method)
			for _, p := range params {
				method.AddArgument(p)
			}
			Expect(method.Arguments.ForMethodSignature()).To(Equal(expectedOutput))
		},
		Entry("there are no arguments", model.Params{}, ""),
		Entry("there is one argument without a name", model.Params{
			{
				Name:       "",
				Type:       "string",
				IsVariadic: false,
			},
		}, "arg0 string"),
		Entry("there are two argument without a name", model.Params{
			{
				Name:       "",
				Type:       "string",
				IsVariadic: false,
			}, {
				Name:       "",
				Type:       "int",
				IsVariadic: false,
			},
		}, "arg0 string, arg1 int"),
		Entry("there are two named argument", model.Params{
			{
				Name:       "str",
				Type:       "string",
				IsVariadic: false,
			}, {
				Name:       "i",
				Type:       "int",
				IsVariadic: false,
			},
		}, "arg0 /* str */ string, arg1 /* i */ int"),
		Entry("there is a single unnamed variadic argument", model.Params{
			{
				Name:       "",
				Type:       "string",
				IsVariadic: true,
			},
		}, "arg0 ...string"),
		Entry("there are multiple unnamed arguments and the last one is variadic", model.Params{
			{
				Name:       "",
				Type:       "int",
				IsVariadic: false,
			}, {
				Name:       "",
				Type:       "string",
				IsVariadic: true,
			},
		}, "arg0 int, arg1 ...string"),
		Entry("there is a single named variadic argument", model.Params{
			{
				Name:       "str",
				Type:       "string",
				IsVariadic: true,
			},
		}, "arg0 /* str */ ...string"),
		Entry("there are multiple unnamed arguments and the last one is variadic", model.Params{
			{
				Name:       "i",
				Type:       "int",
				IsVariadic: false,
			}, {
				Name:       "str",
				Type:       "string",
				IsVariadic: true,
			},
		}, "arg0 /* i */ int, arg1 /* str */ ...string"),
	)
	DescribeTable("ContextVarName",
		func(params model.Params, expectedOutput string) {
			method := new(model.Method)
			for _, p := range params {
				method.AddArgument(p)
			}
			Expect(method.Arguments.ContextVarName()).To(Equal(expectedOutput))
		},
		Entry("there are no arguments", model.Params{}, ""),
		Entry("the first argument is not of type context.Context", model.Params{
			{
				Name:       "",
				Type:       "string",
				IsVariadic: false,
			},
		}, ""),
		Entry("the first and only argument (unnamed) is of type context.Context", model.Params{
			{
				Name:       "",
				Type:       "context.Context",
				IsVariadic: false,
			},
		}, "arg0"),
		Entry("the first and only argument (named) is of type context.Context", model.Params{
			{
				Name:       "ctx",
				Type:       "context.Context",
				IsVariadic: false,
			},
		}, "arg0"),
		Entry("the first of many arguments is not of type context.Context", model.Params{
			{
				Name:       "",
				Type:       "string",
				IsVariadic: false,
			}, {
				Name:       "",
				Type:       "string",
				IsVariadic: false,
			},
		}, ""),
		Entry("the first of many arguments (named) is of type context.Context", model.Params{
			{
				Name:       "ctx",
				Type:       "context.Context",
				IsVariadic: false,
			}, {
				Name:       "",
				Type:       "string",
				IsVariadic: false,
			},
		}, "arg0"),
		Entry("none of many arguments (named) is of type context.Context", model.Params{
			{
				Name:       "s1",
				Type:       "string",
				IsVariadic: false,
			}, {
				Name:       "s2",
				Type:       "string",
				IsVariadic: false,
			},
		}, ""),
		Entry("the second of many arguments (named) is of type context.Context", model.Params{
			{
				Name:       "",
				Type:       "string",
				IsVariadic: false,
			}, {
				Name:       "ctx",
				Type:       "context.Context",
				IsVariadic: false,
			},
		}, ""),
	)
	DescribeTable("ForMethodInvocationWithoutContext",
		func(params model.Params, expectedOutput string) {
			method := new(model.Method)
			for _, p := range params {
				method.AddArgument(p)
			}
			Expect(method.Arguments.ForMethodInvocationWithoutContext()).To(Equal(expectedOutput))
		},
		Entry("there are no arguments", model.Params{}, ""),
		Entry("the first argument is not of type context.Context", model.Params{
			{
				Name:       "",
				Type:       "string",
				IsVariadic: false,
			},
		}, "arg0"),
		Entry("the first and only argument is of type context.Context", model.Params{
			{
				Name:       "",
				Type:       "context.Context",
				IsVariadic: false,
			},
		}, ""),
		Entry("the first of many arguments is not of type context.Context", model.Params{
			{
				Name:       "",
				Type:       "string",
				IsVariadic: false,
			}, {
				Name:       "",
				Type:       "string",
				IsVariadic: false,
			},
		}, "arg0, arg1"),
		Entry("the first of many arguments (named) is of type context.Context", model.Params{
			{
				Name:       "ctx",
				Type:       "context.Context",
				IsVariadic: false,
			}, {
				Name:       "s1",
				Type:       "string",
				IsVariadic: false,
			}, {
				Name:       "s2",
				Type:       "string",
				IsVariadic: false,
			},
		}, "arg1, arg2"),
		Entry("none of many arguments (named) is of type context.Context", model.Params{
			{
				Name:       "s1",
				Type:       "string",
				IsVariadic: false,
			}, {
				Name:       "s2",
				Type:       "string",
				IsVariadic: false,
			},
		}, "arg0, arg1"),
		Entry("the second of many arguments (named) is of type context.Context", model.Params{
			{
				Name:       "",
				Type:       "string",
				IsVariadic: false,
			}, {
				Name:       "ctx",
				Type:       "context.Context",
				IsVariadic: false,
			},
		}, "arg0, arg1"),
	)
})

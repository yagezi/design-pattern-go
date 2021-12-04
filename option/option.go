/*
*@Time        : 2021/12/03
*@Creator     : 
*@File        : option.go
*@Description : golang 设计模式：选项模式
 */

package option

type Options struct {
	required   string
	optionStr1 string
	optionStr2 string
	optionInt1 int
	optionInt2 int
}

var defaultOptions = Options{
	optionStr1: "defaultStr1",
	optionStr2: "defaultStr2",
	optionInt1: 1,
	optionInt2: 2,
}

type FuncOption func(options *Options)

func WithOptionStr1(str1 string) FuncOption {
	return func(options *Options) {
		options.optionStr1 = str1
	}
}

func WithOptionInt1(int1 int) FuncOption {
	return func(options *Options) {
		options.optionInt1 = int1
	}
}

func WithOptionStr2AndInt2(str2 string, int2 int) FuncOption {
	return func(options *Options) {
		options.optionStr2 = str2
		options.optionInt2 = int2
	}
}

func NewOptions(requiredStr string, opts ...FuncOption) *Options {
	options := defaultOptions
	for _, o := range opts {
		o(&options)
	}
	options.required = requiredStr
	return &options
}

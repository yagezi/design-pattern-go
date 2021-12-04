/*
*@Time        : 2021/12/03
*@Creator     : 
*@File        : option_test.go
*@Description :
 */

package option

import (
	"fmt"
	"testing"
)

func Test_Options(t *testing.T) {
	opt1 := NewOptions("requiredStr")
	opt2 := NewOptions("requiredStr", WithOptionStr1("mystr1"))
	opt3 := NewOptions("requiredStr", WithOptionStr2AndInt2("mystr2", 22), WithOptionInt1(11))
	print(opt1)
	print(opt2)
	print(opt3)
}

func print(options *Options) {
	fmt.Println(options.required, options.optionStr1, options.optionStr2, options.optionInt1, options.optionInt2)
}

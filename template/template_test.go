/*
*@Time        : 2021/06/21
*@Creator     : 
*@File        : template_test.go
*@Description : test for 模板方法模式
 */

package template

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	h1 := NewHummer(&HummerH1{})
	h1.Run()

	h2 := NewHummer(&HummerH2{
		alarm: true,
	})
	h2.Run()
}


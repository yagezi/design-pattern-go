/*
*@Time        : 2021/06/11
*@Creator     : 
*@File        : proxy.go
*@Description : 代理模式
 */

package proxy

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRunnerProxy(t *testing.T) {
	Convey("Given two runners", t, func() {
		liLei := NewRunner("LiLei")
		hanMeiMei := NewRunner("HanMeiMei")
		Convey("When employ proxy for Lilei", func() {
			liLeiProxy := NewRunnerProxy(liLei)
			Convey("Then proxy dosen't allow he to run", func() {
				liLeiProxy.Run()
			})
		})
		Convey("When employ proxy for HanMeiMei", func() {
			hanMeiMeiProxy := NewRunnerProxy(hanMeiMei)
			Convey("Then proxy allows she to run", func() {
				hanMeiMeiProxy.Run()
			})
		})

	})
}

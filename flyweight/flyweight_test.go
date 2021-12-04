/*
*@Time        : 2021/06/26
*@Creator     : 
*@File        : flyweight.go
*@Description : 享元模式。使用sync.Map实现
 */

package flyweight

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetInstance(t *testing.T) {
	Convey("sequential", t, func() {
		Convey("Given one color, then get the same ball\n", func() {
			color := "red"
			So(GetInstance(color), ShouldEqual, GetInstance(color))
		})

		Convey("Given two colors, then get different balls\n", func() {
			color := "blue"
			anoherColor := "green"
			So(GetInstance(color), ShouldNotEqual, GetInstance(anoherColor))
		})

	})
}

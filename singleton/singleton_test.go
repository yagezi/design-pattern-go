/*
*@Time        : 2021/06/08
*@Creator     : 
*@File        : singleton.go
*@Description : 线程安全的饿汉式 单例模式
 */

package singleton

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSingleTon(t *testing.T) {
	Convey("Given a Singleton implement with hungry pattern", t, func() {
		Convey("When run get method sequecelly", func() {
			lhs := GetInstance()
			rhs := GetInstance()

			Convey("Then instance should be identical", func() {
				So(lhs, ShouldEqual, rhs)
			})
		})

	})
}

func BenchmarkSingleTon(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Error("test failed")
			}
		}
	})
}

func TestSingleTonLazy(t *testing.T) {
	Convey("Given a Singleton implement with lazy pattern", t, func() {
		So(GetInstanceLazy(), ShouldEqual, GetInstanceLazy())
	})
}

func BenchmarkSingleTonLazy(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstanceLazy() != GetInstanceLazy() {
				b.Error("test failed")
			}
		}
	})
}

func TestSingleTonDoubleCheck(t *testing.T) {
	Convey("Given a Singleton implement with double check", t, func() {
		So(GetInstanceDC(), ShouldEqual, GetInstanceDC())
	})
}

func BenchmarkSingleTonDoubleCheck(b *testing.B) {
	b.SetParallelism(2)
	for i := 0; i < b.N; i++ {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if GetInstanceDC() != GetInstanceDC() {
					b.Errorf("test failed")
				}
			}
		})
		DestroyInstance()
	}
}

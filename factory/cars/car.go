/*
*@Time        : 2021/06/09
*@Creator     : 
*@File        : car.go
*@Description : 工厂模式，抽象产品类
 */
package cars

type Car interface {
	Run()
	Stop()

	IsRun() bool
}

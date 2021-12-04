/*
*@Time        : 2021/06/09
*@Creator     : 
*@File        : standard.go
*@Description : 标准工厂模式
 */

package factory

import "github.com/yagezi/design_pattern/factory/cars"

type AbstractCarFactory interface {
	NewCar(t string) cars.Car
}

type CarFactory struct{}

func (c *CarFactory) NewCar(t string) cars.Car {
	switch t {
	case "ford":
		return &cars.Ford{}
	case "honda":
		return &cars.Honda{}
	}
	return nil
}

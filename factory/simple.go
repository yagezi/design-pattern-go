/*
*@Time        : 2021/06/09
*@Creator     : 
*@File        : simple.go
*@Description : 简单工厂模式
 */

package factory

import "github.com/yagezi/design_pattern/factory/cars"

func NewCar(t string) cars.Car {
	switch t {
	case "ford":
		return &cars.Ford{}
	case "honda":
		return &cars.Honda{}
	}
	return nil
}

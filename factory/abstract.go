/*
*@Time        : 2021/06/09
*@Creator     :
*@File        : abstract.go
*@Description : 抽象工厂模式
 */

package factory

import (
	"github.com/yagezi/design_pattern/factory/cars"
	"github.com/yagezi/design_pattern/factory/telephones"
)

type AbstractCreator interface {
	NewCar(t string) cars.Car
	NewPhone(t string) telephones.Phone
}

type Creator1 struct{}

func (c1 *Creator1) NewCar(t string) cars.Car {
	switch t {
	case "ford":
		return &cars.Ford{}
	case "honda":
		return &cars.Honda{}
	}
	return nil
}

func (c1 *Creator1) NewPhone(t string) telephones.Phone {
	switch t {
	case "iphone":
		return &telephones.Iphone{}
	case "nokia":
		return &telephones.Nokia{}
	}
	return nil
}

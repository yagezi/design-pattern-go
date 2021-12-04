
/*
*@Time        : 2021/06/21
*@Creator     : lu.kaicheng 10261316
*@File        : builder.go
*@Description : 建造者模式
 */

package builder

import (
	"fmt"
	"github.com/modood/table"
)


type CarModelInterFace interface {
	EngineBoom()
	Alarm()
	Stop()
}

// 使用模板方法模式封装通用方法
type CarModel struct {
	CarModelInterFace
	sequence []string
}

func NewCarModel(in CarModelInterFace) *CarModel {
	return &CarModel{CarModelInterFace: in}
}

func (c *CarModel) SetSequence(seq []string) {
	c.sequence = seq
}

func (c *CarModel) Run() {
	// table := map[string]func{
	// 	"EngineBoom" : c.EngineBoom,
	// 	"Alarm": c.Alarm,
	// 	"Stop": c.Stop,
	// }
	table := make(map[string]func())
	table["EngineBoom"] = c.EngineBoom
	table["Alarm"] = c.Alarm
	table["Stop"] = c.Stop

	for _, v range := c.sequence {
		if fun, ok := table[v]; ok{
			fun()
		}
	}
}

type BMWModel struct {
}

func (b BMWModel) Run() {
	fmt.Println("BWM car run")
}

func (b BMWModel) Alarm() {
	fmt.Println("BMW car alarm")
}

func (b BMWModel) Stop() {
	fmt.Println("BWM car stop")
}

type CarBuilder interface {
	SetSequence(seq []string)
	GetCarModel()
}

type BMWCarBuilder struct {
	bmw CarModel
}

func NewBMWCarBuilder() *BMWCarBuilder  {
	return &BMWCarBuilder{bmw: BMWModel{}}
}

func (b *BMWCarBuilder) SetSequence(seq []string) {
	b.bmw.SetSequence(seq)
}

/*
*@Time        : 2021/06/21
*@Creator     : 
*@File        : template.go
*@Description : 模板方法模式
 */

package template

import "fmt"

type HummerInterface interface {
	Start()
	Alarm()
	isAlarm() bool
}

type Hummer struct {
	HummerInterface
}

func NewHummer(hi HummerInterface) *Hummer {
	return &Hummer{HummerInterface: hi}
}

func (hm Hummer) Run() {
	hm.Start()
	if hm.isAlarm() {
		hm.Alarm()
	}
}

type HummerH1 struct {
	Hummer
}

func (h HummerH1) Start() {
	fmt.Println("hummer h1 start ...")
}

func (h HummerH1) Alarm() {
	fmt.Println("hummer h1 alarm ...")
}

// hook function
func (h HummerH1) isAlarm() bool {
	return false
}

type HummerH2 struct {
	Hummer
	alarm bool
}

func (h HummerH2) Start() {
	fmt.Println("hummer h2 start ...")
}

func (h HummerH2) Alarm() {
	fmt.Println("hummer h2 alarm ...")
}

func (h HummerH2) isAlarm() bool {
	return h.alarm
}

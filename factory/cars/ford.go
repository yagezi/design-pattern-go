/*
*@Time        : 2021/06/09
*@Creator     :
*@File        : ford.go
*@Description : 工厂模式，具体产品类
 */
package cars

import (
	"fmt"
	"strings"
)

type Ford struct {
	status string
}

func (f *Ford) Run() {
	fmt.Println("Ford car runs")
	f.status = "run"
}

func (f *Ford) Stop() {
	fmt.Println("Ford car stops")
	f.status = "stop"
}

func (f *Ford) IsRun() bool {
	return strings.Compare(f.status, "run") == 0
}

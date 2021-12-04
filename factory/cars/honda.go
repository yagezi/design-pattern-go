/*
*@Time        : 2021/06/09
*@Creator     : 
*@File        : honda.go
*@Description : 工厂模式，具体产品类
 */

package cars

import (
	"fmt"
	"strings"
)

type Honda struct {
	status string
}

func (f *Honda) Run() {
	fmt.Println("Honda car runs")
	f.status = "run"
}

func (f *Honda) Stop() {
	fmt.Println("Ford car stops")
	f.status = "stop"
}

func (f *Honda) IsRun() bool {
	return strings.Compare(f.status, "run") == 0
}

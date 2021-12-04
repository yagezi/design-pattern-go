/*
*@Time        : 2021/06/11
*@Creator     : 
*@File        : proxy.go
*@Description : 代理模式
 */

package proxy

import (
	"fmt"
	"strings"
)

type IRunner interface {
	Run()
	Name() string
}

type Runner struct {
	IRunner
	name string
}

func NewRunner(_name string) *Runner {
	return &Runner{name: _name}
}

func (r *Runner) Run() {
	fmt.Println("runner star to run ...")
}

func (r *Runner) Name() string {
	return r.name
}

type RunnerProxy struct {
	runner IRunner
}

func NewRunnerProxy(_runner IRunner) *RunnerProxy {
	return &RunnerProxy{runner: _runner}
}

func (rp *RunnerProxy) Run() {
	if strings.Compare(rp.runner.Name(), "LiLei") != 0 {
		rp.runner.Run()
	}
}

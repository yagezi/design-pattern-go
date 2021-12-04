/*
*@Time        : 2021/07/05
*@Creator     : 
*@File        : state.go
*@Description : 状态模式
 */

package state

import (
	"fmt"
	"testing"
)

func TestStateMachine(t *testing.T) {
	machine := NewStateMachine(&StopState{})
	machine.GetState()

	fmt.Println("start 1st")
	machine.Run()
	machine.GetState()

	fmt.Println("start 2nd")
	machine.Run()
	machine.GetState()

	fmt.Println("stop")
	machine.Stop()
	machine.GetState()

}

// test Machine in state2.go
func TestMachine(t *testing.T) {
	machine := NewMachine()
	machine.GetState()
	fmt.Println("id is", machine.GetId())

	machine.Open()
	machine.GetState()
	fmt.Println("id is", machine.GetId())

	machine.Open()
	machine.GetState()
	fmt.Println("id is", machine.GetId())

	machine.Close()
	machine.GetState()
	fmt.Println("id is", machine.GetId())

}

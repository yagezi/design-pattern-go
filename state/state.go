/*
*@Time        : 2021/07/05
*@Creator     : 
*@File        : state.go
*@Description : 状态模式
 */

package state

import "fmt"

type IState interface {
	String() string

	Run(machine *StateMachine)
	Stop(machine *StateMachine)
}

type RunState struct{}

func (r RunState) String() string {
	return "running state"
}

func (r RunState) Run(machine *StateMachine) {
	fmt.Println("do nothing")
}

func (r RunState) Stop(machine *StateMachine) {
	fmt.Println("lift is going to stop")
	machine.SetState(&StopState{})
}

type StopState struct{}

func (s StopState) String() string {
	return "stop status"
}

func (s StopState) Run(machine *StateMachine) {
	fmt.Println("lift is going to run")
	machine.SetState(&RunState{})
}

func (s StopState) Stop(machine *StateMachine) {
	fmt.Println("do nothing")
}

type StateMachine struct {
	state IState
}

func (sm StateMachine) GetState() {
	fmt.Println("current status: ", sm.state)
}

func (sm *StateMachine) SetState(state_ IState) {
	sm.state = state_
}

func (sm *StateMachine) Run() {
	sm.state.Run(sm)
}

func (sm *StateMachine) Stop() {
	sm.state.Stop(sm)
}

func NewStateMachine(state_ IState) *StateMachine {
	return &StateMachine{state: state_}
}

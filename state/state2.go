/*
*@Time        : 2021/07/05
*@Creator     : 
*@File        : state2.go
*@Description : 状态模式。另一种实现：状态类中带状态
 */

package state

import "fmt"

type StateInterface interface {
	String() string

	Open()
	Close()
}

// Open state
type OpenState struct {
	state *Machine
}

func NewOpenState(m *Machine) *OpenState {
	return &OpenState{state: m}
}

func (o OpenState) String() string {
	return "Open state"
}

func (o OpenState) Open() {
	fmt.Println("do nothing")
}

func (o *OpenState) Close() {
	fmt.Println("close machine")
	o.state.SetState(NewCloseState(o.state))
	o.state.SetId(o.state.GetId() - 1)
}

// Close State
type CloseState struct {
	state *Machine
}

func NewCloseState(m *Machine) *CloseState {
	return &CloseState{state: m}
}

func (c CloseState) String() string {
	return "Close state"
}

func (c *CloseState) Open() {
	fmt.Println("open machine")
	c.state.SetState(NewOpenState(c.state))
	c.state.SetId(c.state.GetId() + 1)
}

func (c CloseState) Close() {
	fmt.Println("do nothing")
}

type Machine struct {
	state StateInterface
	id    int
}

func NewMachine() *Machine {
	var ptr *Machine = &Machine{}
	ptr.id = 0
	ptr.state = NewCloseState(ptr)
	return ptr
}

func (m Machine) GetState() {
	fmt.Println("current status: ", m.state)
}

func (m Machine) GetId() int {
	return m.id
}

func (m *Machine) SetState(state_ StateInterface) {
	m.state = state_
}

func (m *Machine) SetId(id_ int) {
	m.id = id_
}

func (m *Machine) Open() {
	m.state.Open()
}

func (m *Machine) Close() {
	m.state.Close()
}

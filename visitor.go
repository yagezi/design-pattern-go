/*
*@Time        : 2021/07/05
*@Creator     : 
*@File        : visitor.go
*@Description : 访问者模式。由于go没有方法重载，使用类型断言代替
 */

package visitor

import (
	"fmt"
)

type Employee interface {
	Accept(visitor IVisitor)
}

type Worker struct {
}

func (w Worker) Echo() {
	fmt.Println("I'm a worker")
}

func (w Worker) Accept(visitor IVisitor) {
	visitor.Visit(w)
}

type Manager struct {
}

func (m Manager) Say() {
	fmt.Println("I'm a manager")
}

func (m Manager) Accept(visitor IVisitor) {
	visitor.Visit(m)
}

type IVisitor interface {
	Visit(employee Employee)
}

type Visitor struct {
}

func (v Visitor) Visit(employee Employee) {
	switch e := employee.(type) {
	case Worker:
		e.Echo()
	case Manager:
		e.Say()
	default:
		fmt.Println("Visit ERROR")
	}
}

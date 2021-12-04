/*
*@Time        : 2021/07/05
*@Creator     : 
*@File        : visitor.go
*@Description : 访问者模式。由于go没有方法重载，使用类型断言代替
 */

package visitor

import "testing"

func TestVisitor_Visit(t *testing.T) {
	visitor := &Visitor{}
	for _, e := range MockEmployee() {
		e.Accept(visitor)
	}

}

func MockEmployee() []Employee {
	return []Employee{
		&Worker{},
		&Manager{},
		&Manager{},
		&Worker{},
	}
}

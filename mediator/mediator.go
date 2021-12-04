/*
*@Time        : 2021/06/22
*@Creator     : 
*@File        : mediator.go
*@Description : 中介模式
*
 */

package mediator

import "fmt"

type Mediator interface {
	ForwardMsg(msg string, country Country)
}

type UnitedNation struct {
	USA
	Iran
}

func (un UnitedNation) ForwardMsg(msg string, country Country) {
	switch country.(type) {
	case USA:
		un.Iran.ReceiveMsg(msg)
	case Iran:
		un.USA.ReceiveMsg(msg)
	default:
		fmt.Println("unknwon country")
	}
}

/*
*@Time        : 2021/06/22
*@Creator     : 
*@File        : mediator.go
*@Description : 中介模式
*
*
 */

package mediator

import (
	"fmt"
	"testing"
)

func TestMediator(t *testing.T) {
	un := &UnitedNation{}
	usa := NewUSA(un)
	iran := NewIran(un)

	usa.SendMsg("交出核武器")
	fmt.Println(iran.GetMsg())

	iran.SendMsg("不交")
	fmt.Println(usa.GetMsg())
}

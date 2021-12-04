/*
*@Time        : 2021/06/24
*@Creator     : 
*@File        : country.go
*@Description : 中介者模式，collegue类
 */

package mediator

import "fmt"

var (
	usaMessage  string
	iranMessage string
)

type Country interface {
	SendMsg(msg string)
	ReceiveMsg(msg string)
	GetMsg() string
}

type USA struct {
	mediator Mediator
}

func NewUSA(m Mediator) *USA {
	return &USA{mediator: m}
}

func (usa USA) SendMsg(msg string) {
	usa.mediator.ForwardMsg(msg, usa)
}

func (usa USA) ReceiveMsg(msg string) {
	usaMessage = msg
}

func (usa USA) GetMsg() string {
	return fmt.Sprintln("USA get message: ", usaMessage)
}

type Iran struct {
	mediator Mediator
}

func NewIran(m Mediator) *Iran {
	return &Iran{mediator: m}
}

func (iran Iran) SendMsg(msg string) {
	iran.mediator.ForwardMsg(msg, iran)
}

func (iran Iran) ReceiveMsg(msg string) {
	iranMessage = msg
}

func (iran Iran) GetMsg() string {
	return fmt.Sprintln("Iran get message: ", iranMessage)
}

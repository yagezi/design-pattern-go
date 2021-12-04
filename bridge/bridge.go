
/*
*@Time        : 2021/06/25
*@Creator     : 
*@File        : bridge.go
*@Description : 桥接模式。抽象和实现解耦
 */

package bridge

import "fmt"

// 实现类接口
type IMsgSender interface {
	Send(msg string)
}

// 具体实现类
type EmailMsgSender struct {
}

func NewEmailMsgSender() *EmailMsgSender {
	return &EmailMsgSender{}
}

func (e EmailMsgSender) Send(msg string) {
	fmt.Println("send email: ", msg)
}

// 抽象类接口
type INotification interface {
	Notify(msg string)
}

// 具体抽象类
type ErrorNotification struct {
	sender IMsgSender
}

func NewErrorNotification(_sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: _sender}
}

func (e ErrorNotification) Notify(msg string) {
	e.sender.Send(msg)
}

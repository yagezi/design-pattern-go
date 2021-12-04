/*
*@Time        : 2021/06/25
*@Creator     : 
*@File        : bridge.go
*@Description : 桥接模式。抽象和实现解耦
 */

package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {
	msgSender := NewEmailMsgSender()
	notificator := NewErrorNotification(msgSender)
	notificator.Notify("hello, world")
}

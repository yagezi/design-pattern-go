/*
*@Time        : 2021/07/03
*@Creator     : 
*@File        : command2.go
*@Description : 命令模式。函数式编程实现
 */

package command

import "fmt"

// Command 类型
type TCommand func() error

func StartCommandFunc(user_ string) TCommand {
	start := &Start{}
	return func() error {
		start.Login(user_)
		return start.Run()
	}
}

func ArchiveCommandFunc() TCommand {
	archive := &Archive{}
	return func() error {
		archive.Save()
		return nil
	}
}

func Invoke(fn TCommand) {
	if err := fn(); err != nil {
		fmt.Println(err.Error())
	}
}

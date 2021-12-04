
/*
*@Time        : 2021/07/03
*@Creator     : lu.kaicheng 10261316
*@File        : command.go
*@Description : 命令模式。面向对象实现
 */

package command

import "fmt"

// Game Start 功能类
type Start struct {
	user string
}

func (s *Start) Login(user_ string) {
	s.user = user_
}

func (s Start) Run() error {
	fmt.Println("game start")
	return nil
}

// Game Archive 功能类
type Archive struct {
}

func (a Archive) Save() {
	fmt.Println("save game data")
}

// Command 接口
type ICommand interface {
	Execute() error
}

// Game Start 命令封装类
type StartCommand struct {
	start Start
	user  string
}

func NewStartCommand(user_ string) *StartCommand {
	return &StartCommand{
		start: Start{},
		user:  user_,
	}
}

func (sc *StartCommand) Execute() error {
	sc.start.Login(sc.user)
	return sc.start.Run()
}

// Game Archive 命令封装类
type ArchiveCommand struct {
	archive Archive
}

func NewArchiveCommand() *ArchiveCommand {
	return &ArchiveCommand{}
}

func (ac ArchiveCommand) Execute() error {
	ac.archive.Save()
	return nil
}

// Invoker 用户调用类
type Invoker struct {
	command ICommand
}

func NewInvoker(command_ ICommand) *Invoker {
	return &Invoker{command: command_}
}

func (invoker Invoker) Action() {
	invoker.command.Execute()
}

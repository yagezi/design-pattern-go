/*
*@Time        : 2021/07/03
*@Creator     : 
*@File        : memento.go
*@Description : 备忘录模式
 */

package memento

type Content string

// 备忘录管理类
type Memento struct {
	content Content
}

func (m Memento) CreateSnapshot() {
}

func (m Memento) Get() {
}

func (m Memento) Append() {
}

func (m Memento) Restore() {
}

type Snapshot struct {
	content Content
}

func (s Snapshot) Get() Content {
	return s.content
}

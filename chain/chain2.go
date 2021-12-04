/*
*@Time        : 2021/07/05
*@Creator     : lu.kaicheng 10261316
*@File        : chain2.go
*@Description : 责任链模式。链式实现
 */

package chain

import "fmt"

// Handler抽象类，使用了模板方法模式
type IHandler interface {
	Do() bool
}

type Handler struct {
	IHandler
	successor *Handler
}

func NewHandler(h IHandler) *Handler {
	return &Handler{IHandler: h}
}

func (h *Handler) Handle() {
	isHandled := h.Do()
	if h.successor != nil && !isHandled {
		h.successor.Handle()
	}
}

func (h *Handler) SetSuccessor(handle_ *Handler) {
	h.successor = handle_
}

// HandleChain 类
type HandleChain struct {
	head, tail *Handler
}

func NewHandleChain() *HandleChain {
	return &HandleChain{}
}

func (hc *HandleChain) AddHandler(handler *Handler) {
	if hc.head == nil {
		hc.head = handler
		hc.tail = handler
		return
	}
	hc.tail.SetSuccessor(handler)
	hc.tail = handler
}

func (hc *HandleChain) Handle() {
	if hc.head != nil {
		hc.head.Handle()
	}
}

// Concrete Hanldler
type HandleA struct {
	Handler
}

func (ha HandleA) Do() bool {
	fmt.Println("handle A")
	return false
}

// Concrete Hanldler
type HandleB struct {
	Handler
}

func (hb HandleB) Do() bool {
	fmt.Println("handle B")
	return true
}

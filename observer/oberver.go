/*
*@Time        : 2021/06/29
*@Creator     : 
*@File        : observer.go
*@Description : 观察者模式 / 发布-订阅模式 / 生产者-消费者模式
*               同步阻塞，单进程。
*				Notice方法中，所有观察者的Update方法全都执行完，被观察者才继续执行。
 */

package observer

import (
	"fmt"
)

type IPublisher interface {
	AddMember(members ...ISubscirbe)
	RemoveMember(members ...ISubscirbe)
	NotifyAll(msg string)
}

type AbstracPublisher struct {
	subscirbes map[ISubscirbe]struct{}
}

func NewAbstracPublisher() *AbstracPublisher {
	return &AbstracPublisher{subscirbes: make(map[ISubscirbe]struct{})}
}

func (a *AbstracPublisher) AddMember(members ...ISubscirbe) {
	for _, m := range members {
		a.subscirbes[m] = struct{}{}
	}
}

func (a *AbstracPublisher) RemoveMember(members ...ISubscirbe) {
	for _, m := range members {
		if _, ok := a.subscirbes[m]; !ok {
			fmt.Println("can not find the subscirbe, ", m)
		}
		delete(a.subscirbes, m)
	}
}

func (a AbstracPublisher) NotifyAll(msg string) {
	for member := range a.subscirbes {
		member.Update(msg)
		// go member.Update(msg) // 异步非阻塞
	}
}

type Publisher struct {
	AbstracPublisher
}

func NewPublisher() *Publisher {
	p := &Publisher{}
	p.subscirbes = make(map[ISubscirbe]struct{})
	return p
}

func (p Publisher) SendMsg(msg string) {
	p.NotifyAll(msg)
}

type ISubscirbe interface {
	Update(msg string)
}

type Subscirbe struct {
	message string
}

func NewSubscribe() *Subscirbe {
	return &Subscirbe{}
}

func (s *Subscirbe) Update(msg string) {
	s.message = msg
}

func (s Subscirbe) GetMsg() string {
	return s.message
}

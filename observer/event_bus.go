/*
*@Time        : 2021/07/07
*@Creator     : 
*@File        : event_bus.go
*@Description : event bus in go。使用反射
 */

package observer

import (
	"fmt"
	"reflect"
	"sync"
)

// Bus 接口
type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

// EventBus 异步事件总线
type EventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{handlers: make(map[string][]reflect.Value)}
}

func (bus *EventBus) Subscirbe(topic string, fn interface{}) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return fmt.Errorf("handler is not function type: %s", v.Kind().String())
	}
	handler, ok := bus.handlers[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	bus.handlers[topic] = handler
	return nil
}

func (bus *EventBus) Publish(topic string, args ...interface{}) {
	fns, ok := bus.handlers[topic]
	if !ok {
		fmt.Println("unregister topic: ", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	// 异步非阻塞
	for i := range fns {
		go fns[i].Call(params)
	}
}

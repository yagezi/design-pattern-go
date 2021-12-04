/*
*@Time        : 2021/06/29
*@Creator     : 
*@File        : observer_test.go
*@Description : test for observer.go
 */

package observer

import (
	"fmt"
	"testing"
	"time"
)

func TestObserver(t *testing.T) {
	publisher := NewPublisher()
	Lilei := NewSubscribe()
	HanMeimei := NewSubscribe()
	Zhang3 := NewSubscribe()

	publisher.AddMember(Lilei, HanMeimei, Zhang3)
	publisher.NotifyAll("First")
	fmt.Println("Li lei get msg: ", Lilei.GetMsg())
	fmt.Println("Han Meimei get msg: ", HanMeimei.GetMsg())
	fmt.Println("Zhang3 get msg: ", Zhang3.GetMsg())

	publisher.RemoveMember(Zhang3)
	publisher.NotifyAll("Second")
	fmt.Println("Li lei get msg: ", Lilei.GetMsg())
	fmt.Println("Han Meimei get msg: ", HanMeimei.GetMsg())
	fmt.Println("Zhang3 get msg: ", Zhang3.GetMsg())

}

func TestEventBus_SingleTopic(t *testing.T) {
	eventBus := NewEventBus()
	eventBus.Subscirbe("alpha", Foo)
	eventBus.Subscirbe("alpha", Foo)
	eventBus.Subscirbe("alpha", Bar)

	eventBus.Publish("alpha", "topic alpha")

	time.Sleep(time.Second * 1)
}

func TestEventBus_NotExistTopic(t *testing.T) {
	eventBus := NewEventBus()
	eventBus.Subscirbe("alpha", Foo)
	eventBus.Subscirbe("alpha", Bar)

	eventBus.Publish("beta", "topic beta")

	time.Sleep(time.Second * 1)
}

func TestEventBus_MultiTopics(t *testing.T) {
	eventBus := NewEventBus()
	eventBus.Subscirbe("alpha", Foo)
	eventBus.Subscirbe("alpha", Bar)
	eventBus.Subscirbe("beta", Foo)
	eventBus.Subscirbe("beta", Bar)

	eventBus.Publish("beta", "topic beta")

	time.Sleep(time.Second * 1)
}

func TestEventBus_NoFunctionType(t *testing.T) {
	var testVal int
	eventBus := NewEventBus()
	err := eventBus.Subscirbe("alpha", testVal)
	fmt.Println("err: ", err)

	eventBus.Publish("alpha", "topic alpha")

	time.Sleep(time.Second * 1)
}

func Foo(ctx string) {
	fmt.Println("Foo msg: ", ctx)
}
func Bar(ctx string) {
	fmt.Println("Bar msg: ", ctx)
}

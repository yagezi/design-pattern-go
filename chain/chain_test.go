/*
*@Time        : 2021/07/05
*@Creator     : 
*@File        : chain.go
*@Description : 责任链模式。列表实现
 */

package chain

import (
	"fmt"
	"testing"
)

func TestFilterChain(t *testing.T) {
	chain := &FilterChain{}
	chain.AddFilter(AdvFilter{}, Political{})

	msg1 := "this is an advertisment"
	msg2 := "this contains political message"
	msg3 := "this is pure message"

	fmt.Println("msg1 has sensor world: ", chain.Run(msg1))
	fmt.Println("msg2 has sensor world: ", chain.Run(msg2))
	fmt.Println("msg3 has sensor world: ", chain.Run(msg3))
}

// test chain2.go
func TestChain2(t *testing.T) {
	chain := NewHandleChain()
	chain.AddHandler(NewHandler(&HandleA{}))
	chain.AddHandler(NewHandler(&HandleA{}))
	chain.AddHandler(NewHandler(&HandleA{}))
	chain.AddHandler(NewHandler(&HandleB{}))
	chain.AddHandler(NewHandler(&HandleA{}))

	chain.Handle()
}

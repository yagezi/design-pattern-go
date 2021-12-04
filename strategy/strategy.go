/*
*@Time        : 2021/06/29
*@Creator     : 
*@File        : strategy.go
*@Description : 策略模式
 */

package strategy

import "fmt"

type IStrategy interface {
	Play()
}

// 长传战术
type Rush struct {
}

func (r Rush) Play() {
	fmt.Println("ready to rush")
}

// 传控战术
type Tikitaka struct {
}

func (t Tikitaka) Play() {
	fmt.Println("ready to tiki-taka")
}

type Strategy struct {
	strategy IStrategy
}

func NewStrategy(s IStrategy) *Strategy {
	return &Strategy{strategy: s}
}

func (s Strategy) PlayGame() {
	fmt.Println("Start to play a soccer match")
	s.strategy.Play()
}

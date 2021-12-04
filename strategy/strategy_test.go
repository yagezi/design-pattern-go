/*
*@Time        : 2021/06/29
*@Creator     : 
*@File        : strategy.go
*@Description : 策略模式
 */

package strategy

import "testing"

func TestStrategy_PlayGame(t *testing.T) {
	coach := NewStrategy(&Tikitaka{})
	coach.PlayGame()
}

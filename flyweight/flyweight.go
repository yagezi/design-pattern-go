/*
*@Time        : 2021/06/26
*@Creator     : 
*@File        : flyweight.go
*@Description : 享元模式。使用sync.Map实现
 */

package flyweight

import (
	"fmt"
	"sync"
)

var ballMap sync.Map

type Ball struct {
	color string
}

func GetInstance(_color string) *Ball {
	ball, found := ballMap.LoadOrStore(_color, &Ball{color: _color})
	if found {
		fmt.Println("find existed ball, color ", _color)
	} else {
		fmt.Println("create new ball, color ", _color)
	}
	return ball.(*Ball)
}

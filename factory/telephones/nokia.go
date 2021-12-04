/*
*@Time        : 2021/06/09
*@Creator     : 
*@File        : nokia.go
*@Description : 具体工厂类
 */

package telephones

import "fmt"

type Nokia struct{}

func (n *Nokia) Call() {
	fmt.Println("nokia 5230 start to call somebody ...")
}




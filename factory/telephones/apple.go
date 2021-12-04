/*
*@Time        : 2021/06/09
*@Creator     : 
*@File        : apple.go
*@Description : 具体产品类
 */

package telephones

import "fmt"

type Iphone struct{}

func (i *Iphone) Call() {
	fmt.Println("iphone begin to call somebody ...")
}

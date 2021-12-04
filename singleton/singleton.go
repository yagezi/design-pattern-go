/*
*@Time        : 2021/06/08
*@Creator     : 
*@File        : singleton.go
*@Description : 线程安全的饿汉式 单例模式
 */

package singleton

type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}

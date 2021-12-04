/*
*@Time        : 2021/06/08
*@Creator     : 
*@File        : singleton_lazy.go
*@Description : 线程安全的懒汉模式
 */

package singleton

import "sync"

type SingletonL struct{}

var (
	singletonL *SingletonL
	once       sync.Once
)

func GetInstanceLazy() *SingletonL {
	once.Do(func() {
		singletonL = &SingletonL{}
	})
	return singletonL
}

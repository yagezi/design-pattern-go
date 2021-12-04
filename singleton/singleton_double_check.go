/*
*@Time        : 2021/06/09
*@Creator     : 
*@File        : singleton_double_check.go
*@Description : 双重检查锁定的单例模式
 */

package singleton

import "sync"

type SingletonDC struct{}

var (
	singletonDc *SingletonDC
	lock        sync.Mutex
)

func GetInstanceDC() *SingletonDC {
	if singletonDc == nil {
		lock.Lock()
		if singletonDc == nil {
			singletonDc = &SingletonDC{}
		}
		lock.Unlock()
	}
	return singletonDc
}

func DestroyInstance() {
	singletonDc = nil
}

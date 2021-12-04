/*
*@Time        : 2021/06/25
*@Creator     : 
*@File        : object_adapter.go
*@Description : 适配器模式扩展：对象适配器，通过组合实现适配
 */

package adapter

type IUser interface {
	Create() error
}

type TecentCloud struct{}

func (tc TecentCloud) Pay() string {
	return "Give me money, you'll be happy"
}

type Azure struct{}

func (az Azure) StartUp() string {
	return "A problem has been detected and Windows has been shut down to prevent damage to your system"
}

// Adapter
type UserAdapter struct {
	tc TecentCloud
	az Azure
}

func NewUserAdapter(tc_ TecentCloud, az_ Azure) *UserAdapter {
	return &UserAdapter{
		tc: tc_,
		az: az_,
	}
}

func (u UserAdapter) Create() error {
	u.tc.Pay()
	u.az.StartUp()
	return nil
}

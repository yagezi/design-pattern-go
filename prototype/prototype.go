/*
*@Time        : 2021/06/21
*@Creator     : 
*@File        : prototype.go
*@Description : 原型模式
 */

package prototype

type Cloneable interface {
	Clone() Cloneable
}

type Plane struct {
	name string
	size int
	m map[string]string
}

func (p Plane) GetName() string {
	return p.name
}

func (p *Plane) SetName(_name string) {
	p.name = _name
}

func (p Plane) GetSize() int {
	return p.size
}

func (p *Plane) SetSize(_size int) {
	p.size = _size
}

func (p Plane) Clone() Cloneable {
	ret := Plane{
		name: p.name,
		size: p.size,
		m: map[string]string{},
	}
	for k,v := range p.m {
		ret.m[k] = v
	}
	return ret
}

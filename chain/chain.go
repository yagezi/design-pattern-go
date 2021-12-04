
/*
*@Time        : 2021/07/05
*@Creator     : lu.kaicheng 10261316
*@File        : chain.go
*@Description : 责任链模式。列表实现
 */

package chain

import "strings"

type IFilter interface {
	Filter(ctx string) bool
}

// 过滤广告,消息中不能含有adv串
type AdvFilter struct{}

func (adv AdvFilter) Filter(ctx string) bool {
	advSensorWorld := "adv"
	return strings.Contains(ctx, advSensorWorld)
}

// 过滤政治敏感，不能含有political串
type Political struct{}

func (po Political) Filter(ctx string) bool {
	politicalSensorWorld := "political"
	return strings.Contains(ctx, politicalSensorWorld)
}

type FilterChain struct {
	filters []IFilter
}

func (fc *FilterChain) AddFilter(filter ...IFilter) {
	fc.filters = append(fc.filters, filter...)
}

// 至少有一个敏感词时，返回True
func (fc FilterChain) Run(ctx string) bool {
	for _, filter := range fc.filters {
		if filter.Filter(ctx) {
			return true
		}
	}
	return false
}

// 执行所有过滤
func (fc FilterChain) RunAll(ctx string) {
	for _, filter := range fc.filters {
		filter.Filter(ctx)
	}
}

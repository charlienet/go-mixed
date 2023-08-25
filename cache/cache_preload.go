package cache

import "context"

// PreLoadItem 预加载数据项
type PreLoadItem struct {
	Key   string
	Value any
}

// PreloadFunc 数据预加载函数定义
type PreloadFunc func(context.Context) ([]PreLoadItem, error)

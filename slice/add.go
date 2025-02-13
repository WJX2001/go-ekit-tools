package slice

import "awesomeProject/internal/slice"

// Add 在index处添加元素
// index 范围应
// 如果index == len(src) 则表示往末尾添加元素

func Add[Src any](src []Src, element Src, index int) ([]Src, error) {
	res, err := slice.Add[Src](src, element, index)
	return res, err
}

package slice

// IntersectSet 取交集，只支持 comparable 类型
// 已去重
func IntersectSet[T comparable](src []T, dst []T) []T {
	srcMap := toMap(src)
	var ret = make([]T, 0, len(src))
	// 交集小于等于两个集合中的任意一个
	for _, val := range dst {
		if _, exist := srcMap[val]; exist {
			ret = append(ret, val)
		}
	}
	return deduplicate[T](ret)
}

func IntersectSetFunc[T any](src []T, dst []T, equal equalFunc[T]) []T {
	var ret = make([]T, 0, len(src))
	for _, val := range dst {
		if ContainsFunc[T](src, func(t T) bool {
			return equal(t, val)
		}) {
			ret = append(ret, val)
		}
	}
	return deduplicateFunc[T](ret, equal)
}

package slice

// Contains 判断 src 里面是否存在 dst
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc(src, func(src T) bool {
		return src == dst
	})
}

// ContainsFunc 判断 src 里面是否存在 dst
// 优先使用 Contains
func ContainsFunc[T any](src []T, equal func(src T) bool) bool {
	// 遍历使用equal 函数进行判断
	for _, v := range src {
		if equal(v) {
			return true
		}
	}
	return false
}

// ContainsAny 判断 src 里面是否存在 dst中的任何一个元素
func ContainsAny[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, exist := srcMap[v]; exist {
			return true
		}
	}
	return false
}

// ContainsAnyFunc 判断 src 里面是否存在 dst 中的任何一个元素
// 应该优先 ContainsAny
func ContainsAnyFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, valDst := range dst {
		for _, valSrc := range src {
			if equal(valSrc, valDst) {
				return true
			}
		}
	}
	return false
}

// ContainsAll 判断 src 里面是否存在 dst 中的所有元素
func ContainsAll[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, exist := srcMap[v]; !exist {
			return false
		}
	}
	return true
}

// ContainsAllFunc 判断 src 里面是否存在 dst中的所有元素
// 首先使用 ContainsAll
func ContainsAllFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, valDst := range dst {
		if !ContainsFunc[T](src, func(src T) bool { return equal(src, valDst) }) {
			return false
		}
	}
	return true
}

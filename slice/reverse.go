package slice

// Reverse 将会完全创建一个新的切片，而不是直接在src上进行翻转
func Reverse[T any](src []T) []T {
	var ret = make([]T, len(src))
	for i := len(src) - 1; i >= 0; i-- {
		ret = append(ret, src[i])
	}
	return ret
}

func ReverseSelf[T any](src []T) {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
}

package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  4,
			name: "dst exist",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  3,
			name: "dst not exist",
		},
		{
			want: false,
			src:  []int{},
			dst:  4,
			name: "length of src is 0",
		},
		{
			want: false,
			dst:  4,
			name: "src nil",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Contains(test.src, test.dst))
		})
	}
}

func TestContainsFunc(t *testing.T) {
	tests := []struct {
		name string
		want bool
		src  []int
		dst  int
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  4,
			name: "dst exist",
		},
		{
			name: "dst not exist",
			src:  []int{1, 4, 6, 2, 6},
			dst:  3,
			want: false,
		},
		{
			want: false,
			src:  []int{},
			dst:  4,
			name: "length of src is 0",
		},
		{
			name: "src nil",
			dst:  4,
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsFunc[int](test.src, func(src int) bool {
				return src == test.dst
			}))
		})
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		name string
		want bool
		src  []int
		dst  []int
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 6},
			name: "exist two ele",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{7, 0},
			name: "not exist the same",
		},
		{
			want: true,
			src:  []int{1, 1, 8},
			dst:  []int{1, 1},
			name: "exist two same ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: false,
			dst:  []int{1},
			name: "src nil",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAny[int](test.src, test.dst))
		})
	}
}

func TestContainsAllFunc(t *testing.T) {
	tests := []struct {
		want bool
		src  []int
		dst  []int
		name string
	}{
		{
			want: true,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2},
			name: "src exist one not in dst",
		},
		{
			want: false,
			src:  []int{1, 4, 6, 2, 6},
			dst:  []int{1, 4, 6, 2, 6, 7},
			name: "src not include the whole ele",
		},
		{
			want: false,
			src:  []int{},
			dst:  []int{1},
			name: "length of src is 0",
		},
		{
			want: true,
			src:  nil,
			dst:  []int{},
			name: "src nil dst empty",
		},
		{
			want: true,
			src:  nil,
			name: "src and dst nil",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, ContainsAllFunc[int](test.src, test.dst, func(src, dst int) bool {
				return src == dst
			}))
		})
	}
}

func ExampleContains() {
	res := Contains[int]([]int{1, 2, 3}, 3)
	fmt.Println(res)
	// Output:
	// true
}

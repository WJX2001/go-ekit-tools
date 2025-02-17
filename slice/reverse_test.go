package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInt(t *testing.T) {
	testCases := []struct {
		name string
		src  []int
		want []int
	}{
		{
			want: []int{7, 5, 3, 1},
			src:  []int{1, 3, 5, 7},
			name: "normal test",
		},
		{
			src:  []int{},
			want: []int{},
			name: "length of src is 0",
		},
		{
			src:  nil,
			want: []int{},
			name: "length of src is nil",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := Reverse[int](tc.src)
			assert.ElementsMatch(t, tc.want, res)
			assert.NotSame(t, tc.src, res)
		})
	}
}

func TestReverseStruct(t *testing.T) {
	type testStruct struct {
		A int
		B []int
	}
	tests := []struct {
		name string
		src  []testStruct
		want []testStruct
	}{
		{
			want: []testStruct{{1, []int{1, 2, 3}}, {3, []int{4, 5, 6}}, {5, []int{7, 8, 9}}},
			src:  []testStruct{{5, []int{7, 8, 9}}, {3, []int{4, 5, 6}}, {1, []int{1, 2, 3}}},
			name: "normal test",
		},
		{
			src:  []testStruct{},
			want: []testStruct{},
			name: "length of src is 0",
		},
		{
			src:  nil,
			want: []testStruct{},
			name: "length of src is nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Reverse[testStruct](tt.src)
			assert.ElementsMatch(t, tt.want, res)
			assert.NotSame(t, tt.src, res)
		})
	}
}

func TestReverseSelfInt(t *testing.T) {
	tests := []struct {
		name string
		src  []int
		want []int
	}{
		{
			want: []int{7, 5, 3, 1},
			src:  []int{1, 3, 5, 7},
			name: "normal test",
		},
		{
			src:  []int{},
			want: []int{},
			name: "length of src is 0",
		},
		{
			src:  nil,
			want: []int{},
			name: "length of src is nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseSelf[int](tt.src)
			assert.ElementsMatch(t, tt.want, tt.src)
		})
	}
}

func TestReverseSelfStruct(t *testing.T) {
	type testStruct struct {
		A int
		B []int
	}
	tests := []struct {
		name string
		src  []testStruct
		want []testStruct
	}{
		{
			want: []testStruct{{1, []int{1, 2, 3}}, {3, []int{4, 5, 6}}, {5, []int{7, 8, 9}}},
			src:  []testStruct{{5, []int{7, 8, 9}}, {3, []int{4, 5, 6}}, {1, []int{1, 2, 3}}},
			name: "normal test",
		},
		{
			src:  []testStruct{},
			want: []testStruct{},
			name: "length of src is 0",
		},
		{
			src:  nil,
			want: []testStruct{},
			name: "length of src is nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseSelf[testStruct](tt.src)
			assert.ElementsMatch(t, tt.want, tt.src)
		})
	}
}

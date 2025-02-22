package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func equal(nums []int, m map[int]struct{}) bool {
	for _, num := range nums {
		_, ok := m[num]
		if !ok {
			return false
		}
		delete(m, num)
	}
	return len(m) == 0
}

func TestSetx_Add(t *testing.T) {
	AddVals := []int{1, 2, 3, 1}
	s := NewMapSet[int](10)
	t.Run("Add", func(t *testing.T) {
		for _, v := range AddVals {
			s.Add(v)
		}
		assert.Equal(t, s.m, map[int]struct{}{
			1: struct{}{},
			2: struct{}{},
			3: struct{}{},
		})
	})
}

func TestSetx_Delete(t *testing.T) {
	testCases := []struct {
		name    string
		delVal  int
		setSet  map[int]struct{}
		wantSet map[int]struct{}
		isExist bool
	}{
		{
			name:   "delete val ",
			delVal: 2,
			setSet: map[int]struct{}{
				2: struct{}{},
			},
			wantSet: map[int]struct{}{},
			isExist: true,
		},
		{
			name:   "deleted val not found",
			delVal: 3,
			setSet: map[int]struct{}{
				2: struct{}{},
			},
			wantSet: map[int]struct{}{
				2: struct{}{},
			},
			isExist: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewMapSet[int](10)
			s.m = tc.setSet
			s.Delete(tc.delVal)
			assert.Equal(t, tc.wantSet, s.m)
		})
	}
}

func TestSetx_IsExist(t *testing.T) {
	s := NewMapSet[int](10)
	s.Add(1)
	testCases := []struct {
		name    string
		val     int
		isExist bool
	}{
		{},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ok := s.Exist(tc.val)
			assert.Equal(t, tc.isExist, ok)
		})
	}
}

func TestSetx_Values(t *testing.T) {
	s := NewMapSet[int](10)
	testCases := []struct {
		name    string
		setSet  map[int]struct{}
		wantVal map[int]struct{}
	}{
		{
			name: "found values",
			setSet: map[int]struct{}{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
			},
			wantVal: map[int]struct{}{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s.m = tc.setSet
			vals := s.Keys()
			ok := equal(vals, tc.wantVal)
			assert.Equal(t, true, ok)
		})
	}
}

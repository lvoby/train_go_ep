package slices

import (
	"testing"
)

func TestMaxMainCases(t *testing.T) {
	var tests = []struct {
		arr       []string
		maxString string
	}{
		{[]string{"a", "bb", "ccc"}, "ccc"},
		{[]string{"a", "b", "c"}, "a"},
		{[]string{"aaa", "b"}, "aaa"},
		{[]string{""}, ""},
	}
	for _, c := range tests {
		got := Max(c.arr)
		if got != c.maxString {
			t.Errorf("Max(%v) == %s, want %s", c.arr, got, c.maxString)
		}
	}
}

func TestMaxEmptyArray(t *testing.T) {
    defer func() {
        if r := recover(); r != nil {
            // it's ok
        }
    }()
    Max([]string{})
    t.Errorf("The code did not panic")
}

func TestReverse(t *testing.T) {
	var tests = []struct {
		arr, reverseArr []int64
	}{
		{[]int64{},[]int64{}},
		{[]int64{1, 2, 3, 4, 5}, []int64{5, 4, 3, 2, 1}},
		{[]int64{1, 1, 1}, []int64{1, 1, 1}},
	}
	for _, c := range tests {
		gotarr := Reverse(c.arr)
		if !equalSlice(gotarr, c.reverseArr) {
			t.Errorf("Reverse(%v) == %v, want %v", c.arr, gotarr, c.reverseArr)
		}
	}
}

func equalSlice(firstSlice, secondSlice []int64) bool {
	if len(firstSlice) != len(secondSlice) {
		return false
	}
	for i := 0; i < len(firstSlice); i ++ {
		if firstSlice[i] != secondSlice[i] {
			return false
		}
	}
	return true
}

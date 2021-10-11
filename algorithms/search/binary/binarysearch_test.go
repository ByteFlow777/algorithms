package search

import (
	"math"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{-2, -1, 0, 1, 2, 3, 4, 5, 6}

	// test for nil slice
	if idx := BinarySearch(nil, math.MaxInt); idx != -1 {
		t.Errorf("result is not right")
	}

	// test for not exist num
	if idx := BinarySearch(nums, math.MaxInt); idx != -1 {
		t.Errorf("result is not right")
	}

	// test for first num
	if idx := BinarySearch(nums, nums[0]); idx != 0 {
		t.Errorf("result is not right")
	}

	// test for mid num
	if idx := BinarySearch(nums, nums[len(nums)/2-1]); idx != len(nums)/2-1 {
		t.Errorf("result is not right")
	}

	// test for last num
	if idx := BinarySearch(nums, nums[len(nums)-1]); idx != len(nums)-1 {
		t.Errorf("result is not right")
	}
}

package search

func BinarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if target < nums[mid] {
			r = mid - 1
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

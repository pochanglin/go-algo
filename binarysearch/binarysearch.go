package binarysearch

// Given an array sorted in asc order and a int target
// search target in array nums, if exist return index, otherwise return -1
// complexity should be O(log n)
func Search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

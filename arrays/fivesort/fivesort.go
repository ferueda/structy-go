package fivesort

func fiveSort(nums []int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[right] == 5 {
			right -= 1
			continue
		}

		if nums[left] == 5 {
			nums[left], nums[right] = nums[right], nums[left]
			right -= 1
		}

		left += 1
	}

	return nums
}

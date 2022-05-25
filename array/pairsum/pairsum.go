package pairsum

func pairSum(nums []int, target int) [2]int {
	seen := make(map[int]int)

	for i, n := range nums {
		comp := target - n

		if v, ok := seen[comp]; ok {
			return [2]int{v, i}
		}

		seen[n] = i
	}

	return [2]int{}
}

package intersection

func intersection(a, b []int) []int {
	var r []int
	set := make(map[int]bool)

	for _, num := range b {
		set[num] = true
	}

	for _, num := range a {
		if _, ok := set[num]; ok {
			r = append(r, num)
		}
	}

	return r
}

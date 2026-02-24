package contains_duplicate

func ContainsDuplicate(nums []int) bool {
	dupMap := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := dupMap[num]; ok {
			return true
		}
		dupMap[num] = struct{}{}
	}
	return false
}

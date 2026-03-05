package happy_number

import (
	"strconv"
)

// IsHappy todo rewrite to Floyd Cycle Detection
// https://leetcode.com/problems/happy-number
func IsHappy(n int) bool {
	_strN := strconv.Itoa(n)
	nums := make(map[int]struct{})
	return _isHappy(_strN, &nums)
}

func _isHappy(n string, nums *map[int]struct{}) bool {
	var res int
	for _, char := range n {
		num, _ := strconv.Atoi(string(char))
		res += num * num
	}
	if _, ok := (*nums)[res]; ok {
		return false
	}
	(*nums)[res] = struct{}{}
	strRes := strconv.Itoa(res)
	if strRes != "1" {
		return _isHappy(strRes, nums)
	}
	return true
}

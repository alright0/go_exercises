package max_sum_subarray

import (
	"math"
)

func MaxSumSubarray(arr []int, k int) int {
	if len(arr) < k {
		return 0
	}

	maxSum := math.MinInt
	var curSum int
	var left int

	for right := 0; right < len(arr); right++ {
		if right-left == k {
			curSum -= arr[left]
			left++
		}
		curSum += arr[right]
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}

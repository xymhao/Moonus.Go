package array

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i, j := 0, 1; j < len(nums); i++ {
		if nums[i] == nums[j] {
			return true
		}
		j++
	}
	return false
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func maxSubArray2(nums []int) int {
	max, sum := -1<<32, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if max < sum {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

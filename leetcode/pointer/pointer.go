package pointer

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/
// 双指针
func sortedSquares(nums []int) []int {
	i := 0
	j := len(nums) - 1
	result := make([]int, len(nums))
	insert := j
	for i <= j {
		left := sqr(nums[i])
		right := sqr(nums[j])
		if left > right {
			result[insert] = left
			insert--
			i++
		} else {
			result[insert] = right
			insert--
			j--
		}
	}
	return result
}

func sqr(val int) int {
	return val * val
}

// https://leetcode-cn.com/problems/rotate-array/
func rotate(nums []int, k int) {
	result := make([]int, len(nums))
	n := len(nums)
	for i, num := range nums {
		index := (i + k) % n
		result[index] = num
	}
	copy(nums, result)
}

// 环状替换
func rotate2(nums []int, k int) {
	n := len(nums)
	for start, count := 0, 0; count < n; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
			count++
		}
	}
}

// 最大公约数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func rotate3(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func moveZeroes(nums []int) {
	left, right := 0, 0
	for right < len(nums) {
		if left == right && nums[left] == 0 {
			right++
			continue
		}

		if left == right && nums[left] != 0 {
			left++
			right++
			continue
		}

		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/submissions/
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum > target {
			right--
		}

		if sum < target {
			left++
		}
	}
	return nil
}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

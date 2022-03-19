package structure

func search(nums []int, target int) int {
	if nums[0] == target {
		return 0
	}

	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}

	left := 0
	right := len(nums)
	for {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if left == right || left == mid || right == mid {
			return -1
		}

		if nums[mid] > target {
			right = mid
			continue
		}

		left = mid
	}
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left := 1
	right := n
	for true {
		mid := (left + right) / 2
		if right-left <= 1 {
			if isBadVersion(left) {
				return left
			}
			return right
		}
		if !isBadVersion(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	return 0
}

func isBadVersion(version int) bool {
	if version >= 1 {
		return true
	}
	return false
}

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	if nums[left] >= target {
		if nums[left] < target {
			return 0
		}
		return left
	}

	if nums[right] <= target {
		if target > nums[right] {
			return right + 1
		}
		return right
	}

	for true {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if right-left <= 1 {
			if nums[right] < target {
				right++
			}
			return right
		}

		if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	return 0
}

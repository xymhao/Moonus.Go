package structure

//https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, num := range nums {
		if j, ok := hashTable[target-num]; ok {
			return []int{j, i}
		}
		hashTable[num] = i
	}
	return nil
}

//https://leetcode-cn.com/problems/merge-sorted-array/

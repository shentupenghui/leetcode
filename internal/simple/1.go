package simple

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for index, num := range nums {
		if i, ok := hashTable[target-num]; ok {
			return []int{i, index}
		}
		hashTable[num] = index
	}
	return nil
}

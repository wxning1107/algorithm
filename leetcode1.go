package algorithm

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if n, ok := m[target-num]; ok {
			return []int{n, i}
		}
		m[num] = i
	}
	return nil
}

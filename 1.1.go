func twoSum(nums []int, target int) []int {

	// Construct a hashmap with the key as the value and value as the index.
    diffMap := make(map[int]int)
    
    for idx, val := range nums {
        diffMap[target - val] = idx
    }
    
    // Iterate through the slice and get the difference straight from the hashmap and if exist return both indices
    for idx1, val := range nums {
        if idx2, ok := diffMap[val]; ok {
            if idx1 != idx2 {
                return []int {idx1, idx2}
            }
        }
    }
    
    return []int{}
    
}

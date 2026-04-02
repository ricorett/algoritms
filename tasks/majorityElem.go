func majorityElement(nums []int) int {
    mapCount := make(map[int]int)
    for i := 0; i < len(nums); i++{
        mapCount[nums[i]]++
    }
    maxCount := 0
    result := 0

for k, v := range mapCount {
    if v > maxCount {
        maxCount = v
        result = k
    }
}

return result
    
}

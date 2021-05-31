package elementcheck

// RemoveElement valと一致しない値に置き換える
func RemoveElement(nums []int, val int) int {
	i := 0
	lengs := len(nums)
	for j := 0; j < lengs; j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

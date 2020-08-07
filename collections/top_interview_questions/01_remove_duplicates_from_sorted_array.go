package questions

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	l := 0
	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[l] {
			l++
			nums[l] = nums[r]
		}
	}
	return l + 1
}

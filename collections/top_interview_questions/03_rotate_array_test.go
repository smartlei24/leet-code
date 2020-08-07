package questions

func rotate(nums []int, k int) {
	lens := len(nums)
	k = k % lens

	for i, j := 0, lens-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := k, lens-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func rotate1(nums []int, k int) {
	lens := len(nums)
	k = k % lens

	tempArray := make([]int, k)
	for i := 0; i < k; i++ {
		tempArray[i] = nums[lens-k+i]
	}
	for i := 0; i < lens-k; i++ {
		nums[lens-i-1] = nums[lens-i-k-1]
	}
	for i := 0; i < k; i++ {
		nums[i] = tempArray[i]
	}
}

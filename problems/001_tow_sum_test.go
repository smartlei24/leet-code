package problems

import (
	"reflect"
	"testing"
)

func twoSum(nums []int, target int) []int {
	l := len(nums)
	m := make(map[int]int, l-1)
	for i := 0; i < l; i++ {
		c := nums[i]
		if id, ok := m[target-c]; ok {
			return []int{id, i}
		}
		m[c] = i
	}
	return []int{}
}

func TestTowSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	actual := twoSum(nums, 9)
	if !reflect.DeepEqual([]int{0, 1}, actual) {
		t.Errorf("actual is %v", actual)
	}
}

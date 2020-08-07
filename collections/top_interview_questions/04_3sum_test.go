package questions

import (
	"log"
	"reflect"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	lens := len(nums)

	// soft
	sort.Ints(nums)

	res := make([][]int, 0)

	for i := 0; i < lens-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, lens-1
		sum := 0
		for l < r {
			sum = nums[l] + nums[r] + nums[i]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}

	return res
}

func TestThreeSum(t *testing.T) {
	input := []int{0, 0, 0}
	expect := [][]int{{0, 0, 0}}
	actual := threeSum(input)
	if !reflect.DeepEqual(expect, actual) {
		log.Fatalf("compire failed, actual is %v", actual)
	}
}

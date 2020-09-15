package problems

import (
	"math"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var longNums []int
	var shortNums []int
	if len(nums1) < len(nums2) {
		shortNums, longNums = nums1, nums2
	} else {
		shortNums, longNums = nums2, nums1
	}
	m, n := len(shortNums), len(longNums)
	iMin, iMax, halfLen := 0, len(shortNums), (m+n+1)/2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && longNums[j-1] > shortNums[i] {
			iMin = i + 1
		} else if i > iMin && shortNums[i-1] > longNums[j] {
			iMax = i - 1
		} else {
			leftMax := 0.0
			if i == 0 {
				leftMax = float64(longNums[j-1])
			} else if j == 0 {
				leftMax = float64(shortNums[i-1])
			} else {
				leftMax = math.Max(float64(longNums[j-1]), float64(shortNums[i-1]))
			}
			if (m+n)%2 == 1 {
				return leftMax
			}

			rightMin := 0.0
			if i == m {
				rightMin = float64(longNums[j])
			} else if j == n {
				rightMin = float64(shortNums[i])
			} else {
				rightMin = math.Min(float64(longNums[j]), float64(shortNums[i]))
			}

			return (leftMax + rightMin) / 2
		}
	}
	return 0
}

func BenchmarkFindMedianSortedArrays(b *testing.B) {
	nums1 := []int{1, 2, 3, 5, 9, 10, 20, 31, 41, 51, 67, 73, 74, 75}
	nums2 := []int{3, 4, 7, 9, 12, 15, 18, 21, 34, 54, 65, 86, 94, 100, 123}
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays(nums1, nums2)
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{3}
	nums2 := []int{1, 2, 4}
	actual := findMedianSortedArrays(nums1, nums2)
	if actual != 2.5 {
		t.Fatalf("actual is %f", actual)
	}
}

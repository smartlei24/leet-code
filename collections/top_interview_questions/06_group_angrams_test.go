package questions

import (
	"reflect"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	t := make(map[string][]string)

	for _, s := range strs {
		runes := []rune(s)
		for i := 0; i < len(runes)-1; i++ {
			for j := i; j < len(runes); j++ {
				if runes[i] < runes[j] {
					runes[i], runes[j] = runes[j], runes[i]
				}
			}
		}
		orderString := string(runes)
		if t[orderString] == nil {
			t[orderString] = make([]string, 0)
		}
		t[orderString] = append(t[orderString], s)
	}

	res := make([][]string, 0, len(t))
	for _, arr := range t {
		res = append(res, arr)
	}
	return res
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expect := [][]string{
		{"ate", "eat", "tea"},
		{"nat", "tan"},
		{"bat"},
	}
	actual := groupAnagrams(strs)
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("actual is %v", actual)
	}
}

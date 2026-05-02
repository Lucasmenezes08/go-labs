package algorithms

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	expected := []int{0, 1}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("esperado %v , resultado %v , dado %v", expected, result, nums)
	}
}

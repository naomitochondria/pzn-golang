package main

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) (int, int) {
	map1 := make(map[int]int)

	for ix, n := range nums {
		if foundIndex, isExist := map1[n]; isExist {
			return foundIndex, ix
		} else {
			map1[target-n] = ix
		}
	}

	return 0, 0
}

func twoSumWithNamedReturn(nums []int, target int) (firstIndex int, secondIndex int) {
	map1 := make(map[int]int)

	for ix, n := range nums {
		if foundIndex, isExist := map1[n]; isExist {
			firstIndex = foundIndex
			secondIndex = ix
		} else {
			map1[target-n] = ix
		}
	}

	return
}

func TestFunction(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 17

	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSumWithNamedReturn(nums, target))
}

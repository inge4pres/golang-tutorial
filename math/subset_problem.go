package subset

import (
	"fmt"
)

func SubsetSolver(input []int, target int) {
	i := Ite(input)
	for sign := i(); sign != 0; sign = i() {
		if sum(arr[i]) == target {
			fmt.Printf("%v\n", arr)
		}
	}
}

func sum(nums []int) (ret int) {
	for i := range nums {
		ret += nums[i]
	}
	return ret
}

func permutations(input []int) (ret [][]int) {
	for length := len(input); length > 1; length-- {
		app := make([]int, lentgh)
		for item := range input {
			app = append(app, input[])
		}
	}
	return ret
}

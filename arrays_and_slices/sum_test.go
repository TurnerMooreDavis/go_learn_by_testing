package arrays_and_slices

import (
	"go-by-example/utils"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("add 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		utils.Assert(t, got, want)
	})
	t.Run("add any number of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		utils.Assert(t, got, want)
		numbers2 := []int{1, 2, 3, 25, 37, 102, 17}
		got2 := Sum(numbers2)
		want2 := 187
		utils.Assert(t, got2, want2)
	})

}

func TestSumAll(t *testing.T) {
	t.Run("Sum multiple slices and return a slice with the sum of each", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{2})
		want := []int{3, 2}
		utils.Assert(t, got, want)
	})
}

func TestSumTails(t *testing.T) {
	t.Run("Sum multiple slices and return a slice with the sum of each without the first number", func(t *testing.T) {
		got := SumTails([]int{1, 2}, []int{2}, []int{})
		want := []int{2, 0, 0}
		utils.Assert(t, got, want)
	})
}

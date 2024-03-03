package arraysslices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of 5 numbers using slices", func(t *testing.T) {
		// numbers := [5]int{1, 2, 3, 4, 5} <- this is an array, it has fixed length
		numbers := []int{1, 2, 3, 4, 5} // <- this is a slice, it has no fixed length

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("expected %d but got %d instead, given %v", expected, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	slice_1 := []int{1, 2}
	slice_2 := []int{0, 9}

	got := SumAll(slice_1, slice_2)
	expected := []int{3, 9}

	// This works but is not type safe, can compare string and slice of int and will compile
	// if !reflect.DeepEqual(expected, got) {
	// 	t.Errorf("Expected %v but got %v instead, given %v, %v", expected, got, slice_1, slice_2)
	// }

	// This works and is type safe but expects the elements to be comparable
	if !slices.Equal(expected, got) {
		t.Errorf("Expected %v but got %v instead, given %v, %v", expected, got, slice_1, slice_2)
	}
}

func TestSumAllTails(t *testing.T) {
	slice_1 := []int{1, 2, 3, 4, 5, 6}
	slice_2 := []int{0, 9, 1, 1, 1, 2}

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails(slice_1, slice_2)
		expected := []int{18, 5}
		checkSums(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails(slice_1, []int{})
		expected := []int{18, 0}
		checkSums(t, got, expected)
	})
}

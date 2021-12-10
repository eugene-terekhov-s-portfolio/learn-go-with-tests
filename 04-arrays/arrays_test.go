package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectSum := func(t testing.TB, actual, expected int) {
		t.Helper()
		if actual != expected {
			t.Errorf("Expected '%d' but got '%d'", expected, actual)
		}
	}

	t.Run("should sum numbers from 1 to 10", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		assertCorrectSum(t, Sum(numbers), 55)
	})
}

func TestSumAllTails(t *testing.T) {
	assertCorrectTails := func(t testing.TB, actual, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected '%v' but got '%v", expected, actual)
		}
	}

	t.Run("should sum all tails", func(t *testing.T) {
		assertCorrectTails(t, SumAllTails([]int{1, 2}, []int{0, 9}), []int{2, 9})
	})

	t.Run("should handle empty slices safely", func(t *testing.T) {
		assertCorrectTails(t, SumAllTails([]int{}, []int{0, 9}), []int{0, 9})
	})
}

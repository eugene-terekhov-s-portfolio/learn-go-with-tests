package arrays

import (
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

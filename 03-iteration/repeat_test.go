package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectRepetition := func(t testing.TB, actual, expected string) {
		t.Helper()
		if actual != expected {
			t.Errorf("Expected %q, but got %q", expected, actual)
		}
	}

	t.Run("Given one symbol, should repeat it 5 times", func(t *testing.T) {
		assertCorrectRepetition(t, Repeat("a", 5), "aaaaa")
	})

	t.Run("Given any symbol and repeat counter, should repeat symbol so many times", func(t *testing.T) {
		assertCorrectRepetition(t, Repeat("x", 3), "xxx")
	})
}

func ExampleRepeat() {
	out := Repeat("y", 6)
	fmt.Println(out)
	// Output: yyyyyy
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

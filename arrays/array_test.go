package arrays

import "testing"

func TestArray(t *testing.T) {

	t.Run("should return sum when a fixed length of array is passed", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := ArraySum(numbers)

		expected := 15

		if got != expected {
			t.Errorf("got %d but expected %d", got, expected)
		}
	})

	t.Run("should return sum when a variable length of array is passed", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := ArraySum(numbers)

		expected := 6

		if got != expected {
			t.Errorf("got %d but expected %d", got, expected)
		}
	})

}

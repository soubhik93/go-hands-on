package arrays

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 3}, []int{4, 5, 6})
	expected := []int{4, 15}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v want %v", got, expected)
	}
}

func TestSumAllTrails(t *testing.T) {

	commonCode := func(t testing.TB, got []int, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v want %v", got, expected)
		}
	}

	t.Run("should add all the tail elements in the array", func(t *testing.T) {
		got := SumAllTrails([]int{1, 3}, []int{4, 5, 6})
		expected := []int{3, 11}

		commonCode(t, got, expected)
	})

	t.Run("should add all the tail elements in the array even if one is empty", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{4, 5, 6})
		expected := []int{0, 11}

		commonCode(t, got, expected)
	})

}

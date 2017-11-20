package sorted_test

import (
	"sorted"
	"testing"
)

// TestStringToFloat64Map for appropriate sorting
func TestStringToFloat64Map(t *testing.T) {
	foo := map[string]float64{
		"hello": 0.98,
		"foo":   -1,
		"world": 0,
		"bar":   100,
	}

	m1 := sorted.NewStringToFloat64Map(foo)
	m2 := sorted.NewStringToFloat64Map(foo)

	m1.Sort()
	m2.SortInDescending()

	keys := make([]string, len(foo))
	values := make([]float64, len(foo))
	i := 0

	m1.Iterate(func(k string, v float64) bool {
		keys[i] = k
		values[i] = v
		i++
		return true
	})

	m2.Iterate(func(k string, v float64) bool {
		i--
		if keys[i] == k && values[i] == v {
			return true
		}

		return false
	})

	if i > 0 {
		t.Fail()
	}
}

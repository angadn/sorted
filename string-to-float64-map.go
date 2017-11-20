package sorted

import (
	"sort"
)

// StringToFloat64Map in a sorted fashion.
type StringToFloat64Map struct {
	data   map[string]float64
	sorted []stringFloat64Pair
}

// NewStringToFloat64Map is a convenience constructor.
func NewStringToFloat64Map(m map[string]float64) StringToFloat64Map {
	return StringToFloat64Map{data: m}
}

// stringFloat64Pair is a key-value pair.
type stringFloat64Pair struct {
	key   string
	value float64
}

// Sort in Ascending.
func (m *StringToFloat64Map) Sort() {
	m.sort(false)
}

// SortInDescending gives the opposite order of Sort.
func (m *StringToFloat64Map) SortInDescending() {
	m.sort(true)
}

func (m *StringToFloat64Map) sort(descending bool) {
	m.sorted = make([]stringFloat64Pair, len(m.data))
	i := 0
	for k, v := range m.data {
		m.sorted[i] = stringFloat64Pair{k, v}
		i++
	}

	sort.Slice(m.sorted, func(i int, j int) bool {
		return descending != (m.sorted[i].value < m.sorted[j].value) // This is an XOR
	})
}

// Iterate in ascending.
func (m StringToFloat64Map) Iterate(cb func(key string, value float64) bool) {
	iterate(m.sorted, cb)
}

func iterate(sorted []stringFloat64Pair, cb func(key string, value float64) bool) {
	for _, pair := range sorted {
		if shouldBreak := !cb(pair.key, pair.value); shouldBreak {
			break
		}
	}
}

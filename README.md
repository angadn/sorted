# sorted
## Data-types supported
###map[string]float64
```
foo := map[string]float64{
	"hello": 0.98,
	"foo":   -1,
	"world": 0,
	"bar":   100,
}

sortedFoo := NewStringToFloat64Map(foo)
sortedFoo.Sort() // or .SortInDescending()
sortedFoo.Iterate(func(k string, v float64) {
	// Everything...in the right place 📻👱 ‍
})
```
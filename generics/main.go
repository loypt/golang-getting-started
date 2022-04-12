package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n\n", SumInts(ints), sumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))
	fmt.Printf("Generic Sums(don't define any type): %v and %v\n\n\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))
	fmt.Printf("Generic Sums with Constraints: %v and %v \n\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsOrFloats[K string, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers[K string, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

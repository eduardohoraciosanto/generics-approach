package main

import "fmt"

var m1 = map[string]int{"first": 10, "second": 20}
var m2 = map[string]float64{"first": 12.34, "second": 56.78}

//SumIntMap returns the sum of all values in the map. This is a non-generic function
func SumIntMap(m map[string]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

//SumFloatMap returns the sum of all values in the map. This is a non-generic function
func SumFloatMap(m map[string]float64) float64 {
	sum := 0.0
	for _, v := range m {
		sum += v
	}
	return sum
}

type Number interface {
	int | float64
}

//SumMapIntOrFloat returns the sum of all values in the map. This is a generic func
func SumMapIntOrFloat[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("Non Generic Int=> ", SumIntMap(m1))
	fmt.Println("Non Generic float=> ", SumFloatMap(m2))

	fmt.Println("Generic Int=> ", SumMapIntOrFloat(m1))
	fmt.Println("Generic float=> ", SumMapIntOrFloat(m2))

}

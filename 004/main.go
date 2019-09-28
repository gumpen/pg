package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)

	fmt.Println(Sqrt(n))
}

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		fmt.Println(z * z)
		z -= (z*z - x) / (2 * x)
		fmt.Println(z)
	}
	return z
}

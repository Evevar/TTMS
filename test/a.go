package main

import "math/rand"

func main() {
	a := 1
	for {
		a++
		a *= rand.Intn(10)
		a /= rand.Intn(10) + 1
		a--
	}
}

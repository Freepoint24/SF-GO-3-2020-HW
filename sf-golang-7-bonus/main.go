package main

import (
	"fmt"
	"sf-golang-7-bonus/archi"
	"sf-golang-7-bonus/primes"
)

func main() {
	source := "Yooooo-hoooooo!!!111"
	compressed := archi.Compress(source)
	decompressed := archi.Decompress(compressed)

	fmt.Println(compressed)
	fmt.Printf("%v\n", source)
	fmt.Printf("%v\n", decompressed)
	fmt.Println(source == decompressed)

	fmt.Printf("%v", primes.Eratosthenes(100))
}

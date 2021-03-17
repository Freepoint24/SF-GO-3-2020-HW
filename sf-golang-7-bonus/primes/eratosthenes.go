package primes

import (
	"math"
)

func Eratosthenes(limit uint) (primes []uint) {
	if limit < 2 {
		return []uint{}
	}

	var i, j uint
	composites := make([]bool, limit-1)

	// TODO откуда такое утверждение? https://habr.com/ru/post/468833/
	scanLimit := uint(math.Sqrt(float64(limit + 1)))

	for i = 2; i <= scanLimit; i++ {
		if composites[i-2] == false {
			for j = i * i; j <= limit; j += i {
				composites[j-2] = true
			}
		}
	}

	for i = 2; i <= limit-2; i++ {
		if !composites[i-2] {
			primes = append(primes, i)
		}
	}
	return
}

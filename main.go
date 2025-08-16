package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	result := make([]int, len(numbers))
	var wg sync.WaitGroup
	wg.Add(len(numbers)) // тут используется кол-во горутин столько, сколько чисел в массиве

	for i, n := range numbers {
		go func(index, value int) {
			defer wg.Done()
			result[index] = value * value
		}(i, n) // использовал index чтобы порядок точно был правильным
	}
	fmt.Println(result)
}

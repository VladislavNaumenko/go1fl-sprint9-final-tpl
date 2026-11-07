package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// обработка коренр-кейса
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Int()
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// обработка коренр-кейса
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, value := range data[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// обработка коренр-кейса
	if len(data) == 0 {
		return 0
	}
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxElements := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		startIndex := i * chunkSize
		endIndex := startIndex + chunkSize
		if i == CHUNKS-1 {
			endIndex = len(data)
		}

		go func(piece []int, idx int) {
			defer wg.Done()
			maxElements[idx] = maximum(piece)
		}(data[startIndex:endIndex], i)
	}
	wg.Wait()

	return maximum(maxElements)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(slice)
	elapsed := time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(slice)
	elapsed = time.Since(start).Milliseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}


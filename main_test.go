package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, tc := range []struct {
		name string
		size int
	}{
		{"нулевая длина", 0},
		{"отрицательная длина", -5},
		{"один элемент", 1},
		{"обычный размер", 1000},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := generateRandomElements(tc.size)
			if len(got) != tc.size {
				t.Fatalf("длина слайса: ожидалось %d, получено %d", tc.size, len(got))
			}
			for _, v := range got {
				if v < 0 {
					t.Fatalf("найдено отрицательное число: %d", v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	for _, tc := range []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"nil-слайс", nil, 0},
		{"один элемент", []int{42}, 42},
		{"все равны", []int{7, 7, 7, 7}, 7},
		{"максимум в начале", []int{99, 1, 2, 3}, 99},
		{"максимум в конце", []int{1, 2, 3, 99}, 99},
		{"максимум в середине", []int{1, 99, 2}, 99},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := maximum(tc.data); got != tc.want {
				t.Fatalf("maximum(%v): ожидалось %d, получено %d", tc.data, tc.want, got)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	for _, tc := range []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{5}, 5},
		{"меньше CHUNKS", []int{3, 9, 1}, 9},
		{"не делится на CHUNKS нацело", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{"ровно CHUNKS", []int{8, 3, 5, 1, 9, 2, 7, 4}, 9},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := maxChunks(tc.data); got != tc.want {
				t.Fatalf("maxChunks(%v): ожидалось %d, получено %d", tc.data, tc.want, got)
			}
		})
	}
}


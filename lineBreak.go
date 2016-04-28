package main

import (
	"fmt"
	"math"
	"strings"
)

// reverse takes a []string and returns a reversed []string
func reverse(slice []string) {
	for left, right := 0, len(slice)-1; left < right; left, right = left+1, right-1 {
		slice[left], slice[right] = slice[right], slice[left]
	}
}

// lineBreak takes a string and a line width,
// and returns a []string of lines optimally broken
func lineBreak(text string, width int) []string {
	words := strings.Split(text, " ")
	count := len(words)
	offsets := []int{0}
	for _, w := range words {
		offsets = append(offsets, offsets[len(offsets)-1]+len(w))
	}

	minima := make([]float64, count+1)
	for x := 1; x < len(minima); x++ {
		minima[x] = math.Pow(10, 20)
	}
	breaks := make([]int, count+1)
	for i := 0; i < count; i++ {
		for j := i + 1; j <= count; j++ {
			w := offsets[j] - offsets[i] + j - i - 1
			if w > width {
				break
			}
			cost := minima[i] + math.Pow(float64(width-w), 2)
			if cost < minima[j] {
				minima[j] = cost
				breaks[j] = i
			}
		}
	}
	lines := []string{}
	j := count
	for j > 0 {
		i := breaks[j]
		lines = append(lines, strings.Join(words[i:j], " "))
		j = i
	}
	reverse(lines)
	return lines
}

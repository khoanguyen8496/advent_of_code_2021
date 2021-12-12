package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// main ...
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prev := -1
	count := 0
	window := make([]int, 0, 3)
	for scanner.Scan() {
		line := scanner.Text()
		current, _ := strconv.Atoi(line)
		window = append(window, current)
		for len(window) > 3 {
			window = window[1:]
		}
		if len(window) < 3 {
			continue
		}
		windowSum := 0
		for i := 0; i < len(window); i++ {
			windowSum += window[i]
		}
		fmt.Printf("prev=%v, current=%v\n", prev, windowSum)
		if prev != -1 && windowSum > prev {
			count++
		}
		prev = windowSum
	}
	fmt.Printf("count is %v\n", count)
}

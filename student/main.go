package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"guessit1/student/mathskills"
)



func main() {
	var dataSlice []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input!!!!")
			continue
		}
		dataSlice = append(dataSlice, num)

		if len(dataSlice) > 1 {
			lower, upper := mathskills.Range(dataSlice)
			fmt.Printf("%d %d\n", lower, upper)
		}
	}
}
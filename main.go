package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Open the text file and read its contents into a string
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Trim the whitespace before calling Split
	contents := strings.TrimSpace(string(data))
	var numbers []float64
	for _, s := range strings.Split(contents, "\n") {
		// Filter out empty strings
		if s == "" {
			continue
		}
		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		numbers = append(numbers, n)
	}

	// Sort the slice of integers
	sort.Float64s(numbers)

	// Calculate the average
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	average := float64(sum) / float64(len(numbers))
	roundedAverage := math.Round(average)

	// Calculate the median
	var median float64
	if len(numbers)%2 == 0 {
		// If the number of values is even, the median is the average of the two middle values
		median = (numbers[len(numbers)/2-1] + numbers[len(numbers)/2]) / 2
	} else {
		// If the number of values is odd, the median is the middle value
		median = (numbers[len(numbers)/2])
	}
	roundedMedian := math.Round(median)

	// Calculate the variance
	var variance float64
	for _, num := range numbers {
		variance += (float64(num) - average) * (float64(num) - average)
	}
	variance /= float64(len(numbers))
	roundedVariance := int(math.Round(variance))

	// Calculate the standard deviation
	stddev := math.Sqrt(variance)
	roundedStddev := math.Round(stddev)

	fmt.Println("Average:", roundedAverage)
	fmt.Println("Median:", roundedMedian)
	fmt.Println("Variance:", roundedVariance)
	fmt.Println("Standard deviation:", roundedStddev)
}

package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		panic("Invalid number of arguments passed.")
	}
	data := ReadInput(args[0])
	avg := CalculateAverage(data, false)
	median := CalculateMedian(data)
	variance := CalculateVariance(data, avg)
	standardDeviation := math.Sqrt(variance)
	fmt.Printf("Average: %d\nMedian: %d\nVariance: %d\nStandard Deviation: %d", int(avg), int(median), int(variance), int(standardDeviation))
}

func ReadInput(filename string) []int {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println(file)
	data := ConvertFile(file)
	return data
}

func ConvertFile(arr []byte) []int {
	var output []int
	isNewNum := false
	current := 0
	for _, b := range arr {
		if 48 <= int(b) && int(b) <= 57 {
			isNewNum = true
			current *= 10
			current += int(b) - 48
		} else {
			if isNewNum {
				output = append(output, current)
			}
			current = 0
			isNewNum = false
		}
	}
	return output
}

func Sum(data []int, squared bool) int {
	sum := 0
	for _, n := range data {
		if squared {
			sum += n * n
		} else {
			sum += n
		}
	}
	return sum
}

func Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)-1]
	var smaller []int
	var bigger []int
	for index, i := range arr {
		if index != len(arr)-1 {
			if i <= pivot {
				smaller = append(smaller, i)
			} else {
				bigger = append(bigger, i)
			}
		}
	}
	output := append(Sort(smaller), pivot)
	output = append(output, Sort(bigger)...)
	return output
}

func CalculateAverage(data []int, squared bool) float64 {
	sum := Sum(data, squared)
	return float64(sum) / float64(len(data))
}

func CalculateMedian(data []int) float64 {
	data = Sort(data)
	if len(data)%2 == 1 {
		return float64(data[len(data)/2])
	}
	return float64(data[(len(data)-1)/2]+data[(len(data)+1)/2]) / 2.0
}

func CalculateVariance(data []int, avg float64) float64 {
	avgSq := CalculateAverage(data, true)
	sqAvg := avg * avg
	return avgSq - sqAvg
}

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// create function that takes int and slice and returns a slice
func CumulativeAverages(inputSlice []float64, averageWindow int) []float64 {
	// slice to hold hold avg
	avg := make([]float64, len(inputSlice))
	runningSum := inputSlice[0]
	divisidor := 1
	avg[0] = inputSlice[0] / float64(divisidor)

	for j := 1; j < len(inputSlice); j++ {
		if j < averageWindow {
			runningSum += inputSlice[j]
			divisidor = divisidor + 1
		} else {
			runningSum -= inputSlice[j-averageWindow]
			runningSum += inputSlice[j]
			divisidor = averageWindow
		}
		fmt.Println(runningSum)
		fmt.Println(float64(divisidor))
		avg[j] = runningSum / float64(divisidor)
	}
	return avg

}

func main() {

	// create a slice with values
	//slice := []float64{0, 1, 2, 3}

	// slice2 := []float64{0, 1, -2, 3, -4, 5, -6, 7, -8, 9}
	// fmt.Println(CumulativeAverages(slice, 3))
	// fmt.Println(CumulativeAverages(slice2, 5))

	// Take command line args - and remove non numerics besides comma and periods
	reg, err := regexp.Compile("[^,0-9,.]+")
	if err != nil {
		log.Fatal(err)
	}

	processedString := reg.ReplaceAllString(os.Args[1], "")

	// Create slice from string - split on comma
	slice3 := strings.Split(processedString, ",")
	length := len(slice3) - 1
	// declare slice of length 
	processSlice := make([]float64, len(slice3)-1)
	 := []float64{}
	sumAverage, err := strconv.Atoi(slice3[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sumAverage)
	for i := 1; i < len(slice3); i++ {
		fmt.Println(slice3[i])
		processSlice[i-1], err = strconv.ParseFloat(slice3[i], 64)
		if err != nil {
			log.Fatal(err)
		}

		processSlice = append(processSlice, res)

	}

	// TODO - split the string - with , and convert to fload slice
	//fmt.Println(CumulativeAverages(processSlice, sumAverage))
}

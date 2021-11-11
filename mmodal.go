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
func CumulativeAverages(inputSlice []float64, intWindow int) []float64 {
	// slice to hold hold avg
	avg := make([]float64, len(inputSlice))
	runningSum := inputSlice[0]
	divisidor := 1
	avg[0] = inputSlice[0] / float64(divisidor)

	for j := 1; j < len(inputSlice); j++ {
		if j < intWindow {
			runningSum += inputSlice[j]
			divisidor = divisidor + 1
		} else {
			runningSum -= inputSlice[j-intWindow]
			runningSum += inputSlice[j]
			divisidor = intWindow
		}
		avg[j] = runningSum / float64(divisidor)
	}
	return avg

}

func main() {

	// In main - take one argument queoted - with numbers - as descibed in assignment
	// like ... #go run mmodal.go "(3, [0, 1, 2, 3])"
	// Take command line args - and remove non numerics besides comma and periods
	reg, err := regexp.Compile("[^,0-9,.]+")
	if err != nil {
		log.Fatal(err)
	}

	processedString := reg.ReplaceAllString(os.Args[1], "")

	// Create slice from string - split on comma
	slice3 := strings.Split(processedString, ",")
	processSlice := make([]float64, len(slice3)-1)

	// get integer window
	intWindow, err := strconv.Atoi(slice3[0])
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < len(slice3); i++ {
		processSlice[i-1], err = strconv.ParseFloat(slice3[i], 64)
		if err != nil {
			log.Fatal(err)
		}

		//processSlice = append(processSlice, res)

	}

	// TODO - split the string - with , and convert to fload slice
	fmt.Println(CumulativeAverages(processSlice, intWindow))
}

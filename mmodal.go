package main

import "fmt"

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
	fmt.Println("vim-go")
	// create a slice with values
	slice := []float64{0, 1, 2, 3}
	slice2 := []float64{0, 1, -2, 3, -4, 5, -6, 7, -8, 9}
	fmt.Println(CumulativeAverages(slice, 3))
	fmt.Println(CumulativeAverages(slice2, 5))

}

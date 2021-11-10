package main

import (
	"testing"
)

func TestMModal(t *testing.T) {

	slice := []float64{0, 1, 2, 3}
	//slice2 := []float64{0, 1, -2, 3, -4, 5, -6, 7, -8, 9}

	sliceRes1 := CumulativeAverages(slice, 3)
	t.Log(sliceRes1)
}

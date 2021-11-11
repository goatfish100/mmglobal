package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMModal1(t *testing.T) {

	slc := []float64{0, 1, 2, 3}
	//slice2 := []float64{0, 1, -2, 3, -4 5, -6, 7, -8, 9}

	sliceRes := CumulativeAverages(slc, 3)
	sliceExp := []float64{0, 0.5, 1, 2}
	// assert that sliceRes is equal to sliceRes
	assert.EqualValues(t, sliceRes, sliceExp)

}

func TestMModal2(t *testing.T) {

	slc := []float64{0, 1, -2, 3, -4, 5, -6, 7, -8, 9}

	sliceRes := CumulativeAverages(slc, 5)
	sliceExp := []float64{0, 0.5, -0.3333333333333333, 0.5, -0.4, 0.6, -0.8, 1, -1.2, 1.4}
	// assert that sliceRes is equal to sliceRes
	assert.EqualValues(t, sliceRes, sliceExp)

}

func TestMModal3(t *testing.T) {

	slc := []float64{1, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100}

	sliceRes := CumulativeAverages(slc, 5)
	sliceExp := []float64{1, 3, 5.333333333333333, 7.75, 10.2, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90}
	// assert that sliceRes is equal to sliceRes
	assert.EqualValues(t, sliceRes, sliceExp)

}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Resource struct {
	//ID          string `json:"id" bson:"_id,omitempty"`
	window int32
	values []float64
}

// Create a resource with a window of 3

// var struc = []Resource{
// 	{3, {1, 2, 3, 4}},
// 	{4, {5, 6, 7, 8}},
// }

func TestMModal(t *testing.T) {

	slice := []float64{0, 1, 2, 3}
	//slice2 := []float64{0, 1, -2, 3, -4, 5, -6, 7, -8, 9}

	sliceRes1 := CumulativeAverages(slice, 3)
	sliceExp := []float64{0, 0.5, 1, 2}
	// assert that sliceRes1 is equal to sliceRes2
	assert.EqualValues(t, sliceRes1, sliceExp)

}

func TestMModal2(t *testing.T) {

	slice := []float64{0, 1, -2, 3, -4, 5, -6, 7, -8, 9}

	//slice2 := []float64{0, 1, -2, 3, -4, 5, -6, 7, -8, 9}

	sliceRes1 := CumulativeAverages(slice, 5)
	sliceExp := []float64{0, 0.5, -0.3333333333333333, 0.5, -0.4, 0.6, -0.8, 1, -1.2, 1.4}
	// assert that sliceRes1 is equal to sliceRes2
	assert.EqualValues(t, sliceRes1, sliceExp)

}

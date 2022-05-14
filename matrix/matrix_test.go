package matrix

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func generateMatrix() *Matrix {
	return New([][]float64{
		{0, 6, 7, 5, 4},
		{3, 2, 8, 0, 3},
		{1, 5, 4, 6, 8},
		{9, 5, 3, 7, 2},
		{0, 2, 0, 3, 6},
	})
}

func TestSwapRows(t *testing.T) {
	matrix := generateMatrix()
	row1 := matrix.GetMatrix()[0]
	row2 := matrix.GetMatrix()[1]
	matrix.SwapRows(0, 1)
	a := assert.New(t)
	a.Equal(row2, matrix.GetMatrix()[0])
	a.Equal(row1, matrix.GetMatrix()[1])
}

func TestFindFistNonZeroColumnRow(t *testing.T) {
	matrix := generateMatrix()
	if index, err := matrix.FindFistNonZeroColumnRow(0, 0, len(matrix.GetMatrix())); err != nil {
		t.Fatal(err)
	} else {
		a := assert.New(t)
		a.Equal(1, index)
	}
}

func TestSetToZeroColumnUpToDown(t *testing.T) {
	matrix := generateMatrix()
	matrix.SwapRows(0, 1)

	for _, el := range matrix.GetMatrix() {
		fmt.Println(el)
	}
	matrix.SetToZeroColumnUpToDown(0)
	for _, el := range matrix.GetMatrix() {
		fmt.Println(el)
	}
	matrix.SetToZeroColumnDownToUp(4)
	for _, el := range matrix.GetMatrix() {
		fmt.Println(el)
	}
}

package matrix

import (
	"CourseWorkParallel/file_processor"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func generateMatrix() *Matrix {
	m := New([][]float64{
		{2, 3, 2, 2},
		{-1, -1, 0, -1},
		{-2, -2, -2, -1},
		{3, 2, 2, 2},
	})
	m.MakeExtended()
	return m
}

func generateMatrixWithFirstZeroRow() *Matrix {
	m := New([][]float64{
		{0, 3, 2, 2},
		{-1, -1, 0, -1},
		{-2, -2, -2, -1},
		{3, 2, 2, 2},
	})

	m.MakeExtended()
	return m
}

func TestMatrix_SwapRows(t *testing.T) {
	matrix := generateMatrix()
	row1 := matrix.GetExtendedMatrix()[0]
	row2 := matrix.GetExtendedMatrix()[1]
	matrix.SwapRows(0, 1)
	a := assert.New(t)
	a.Equal(row2, matrix.GetExtendedMatrix()[0])
	a.Equal(row1, matrix.GetExtendedMatrix()[1])
}

func TestMatrix_FindFistNonZeroColumnRow(t *testing.T) {
	matrix := generateMatrixWithFirstZeroRow()
	if index, err := matrix.FindFistNonZeroColumnRow(0, 0, len(matrix.GetMatrix())); err != nil {
		t.Fatal(err)
	} else {
		a := assert.New(t)
		a.Equal(1, index)
	}
}

func TestMatrix_ProcessRow(t *testing.T) {
	matrix := generateMatrix()
	matrix.ProcessRow(0, 1, 2)
	processedRow := []float64{0, 1, 2, 0, 1, 2, 0, 0}
	a := assert.New(t)
	for i, val := range matrix.GetExtendedMatrix()[1] {
		a.Equal(processedRow[i], val)
	}
	matrix.ProcessRow(0, 2, 2)
	processedRow = []float64{0, 1, 0, 1, 1, 0, 1, 0}
	for i, val := range matrix.GetExtendedMatrix()[2] {
		a.Equal(processedRow[i], val)
	}
	matrix.ProcessRow(0, 3, 2)
	coef := -2.0 / 3.0
	processedRow = []float64{0, coef*2.0 + 3.0, coef*2.0 + 2.0, coef*2.0 + 2.0, 1, 0, 0, coef}
	for i, val := range matrix.GetExtendedMatrix()[3] {
		a.Equal(processedRow[i], val)
	}
}

func TestMatrix_SetToZeroColumnUpToDown(t *testing.T) {
	matrix := generateMatrix()
	matrix.SetToZeroColumnUpToDown(0)
	a := assert.New(t)
	for i := 1; i < matrix.initialSize; i++ {
		a.Equal(0.0, matrix.GetExtendedMatrix()[i][0])
	}
	matrix.SetToZeroColumnUpToDown(1)
	for i := 2; i < matrix.initialSize; i++ {
		a.Equal(0.0, matrix.GetExtendedMatrix()[i][1])
	}
	matrix.SetToZeroColumnUpToDown(2)
	for i := 3; i < matrix.initialSize; i++ {
		a.Equal(0.0, matrix.GetExtendedMatrix()[i][2])
	}
	matrix.SetToZeroColumnUpToDown(3)
	for i := 4; i < matrix.initialSize; i++ {
		a.Equal(0.0, matrix.GetExtendedMatrix()[i][3])
	}
}

func TestMatrix_SetToZeroColumnDownToUp(t *testing.T) {
	matrix := generateMatrix()
	matrix.SetToZeroColumnDownToUp(3)
	a := assert.New(t)
	for i := 2; i >= 0; i-- {
		a.Equal(0.0, matrix.GetExtendedMatrix()[i][3])
	}
	matrix.SetToZeroColumnDownToUp(2)
	for i := 1; i >= 0; i-- {
		a.Equal(0.0, matrix.GetExtendedMatrix()[i][2])
	}
	matrix.SetToZeroColumnDownToUp(1)
	for i := 0; i >= 0; i-- {
		a.Equal(0.0, matrix.GetExtendedMatrix()[i][1])
	}
}

func TestMakeFromFile(t *testing.T) {
	fileProcessor := file_processor.New()
	matrix, err := MakeFromFile("./test_file.csv", fileProcessor)

	if err != nil {
		t.Fatal(err)
	}

	expectedData := [][]float64{
		{1, 2, 3.2},
		{5.5, 4, 3},
		{11, -8, -4.5},
	}

	a := assert.New(t)

	for i, row := range matrix.GetMatrix() {
		for j, val := range row {
			a.Equal(expectedData[i][j], val)
		}
	}
}

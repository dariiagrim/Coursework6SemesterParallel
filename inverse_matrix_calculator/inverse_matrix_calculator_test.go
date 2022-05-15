package inverse_matrix_calculator

import (
	"CourseWorkParallel/file_processor"
	"CourseWorkParallel/matrix"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"os"
	"testing"
)

var calculator InverseMatrixCalculator

func TestMain(m *testing.M) {
	calculator = New()
	os.Exit(m.Run())
}

func generateMatrix() *matrix.Matrix {
	return matrix.New([][]float64{
		{2, 3, 2, 2},
		{-1, -1, 0, -1},
		{-2, -2, -2, -1},
		{3, 2, 2, 2},
	})
}
func TestCalculateInverseMatrixSequential(t *testing.T) {
	m := generateMatrix()
	inverseMatrix, err, ms := calculator.CalculateInverseMatrixSequential(m)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(ms)

	for i, row := range inverseMatrix {
		for j, val := range row {
			inverseMatrix[i][j] = math.Round(val*100) / 100
		}
	}

	checkMatrix := MatrixMultiply(m.GetMatrix(), inverseMatrix)
	identityMatrix := matrix.NewIdentity(m.GetInitialSize())

	a := assert.New(t)

	for i, row := range checkMatrix {
		for j, val := range row {
			a.Equal(identityMatrix.GetMatrix()[i][j], val)
		}
	}

}

func TestTemp(t *testing.T) {
	res := matrix.GenerateRandomMatrixInStringFormat(50)
	fileProcessor := file_processor.New()
	err := fileProcessor.WriteCsv("../matrix.csv", res)

	if err != nil {
		t.Fatal(err)
	}

	m, err := matrix.MakeFromFile("../matrix.csv", fileProcessor)

	if err != nil {
		t.Fatal(err)
	}

	inverseCalculator := New()

	inverse, err, ms := inverseCalculator.CalculateInverseMatrixSequential(m)

	fmt.Println(ms, "seq")

	err = fileProcessor.WriteCsv("../inverse_matrix.csv", matrix.StringifyMatrix(inverse))

	if err != nil {
		t.Fatal(err)
	}

	checkMatrix := MatrixMultiply(m.GetMatrix(), inverse)

	for i, row := range checkMatrix {
		for j, val := range row {
			checkMatrix[i][j] = math.Round(val*10000) / 10000
		}
	}

	err = fileProcessor.WriteCsv("../check_matrix.csv", matrix.StringifyMatrix(checkMatrix))

	if err != nil {
		t.Fatal(err)
	}

	inverseParallel, err, ms := inverseCalculator.CalculateInverseMatrixParallel(m)

	fmt.Println(ms, "parallel")

	err = fileProcessor.WriteCsv("../inverse_matrix_parallel.csv", matrix.StringifyMatrix(inverseParallel))

	if err != nil {
		t.Fatal(err)
	}

	checkMatrixParallel := MatrixMultiply(m.GetMatrix(), inverse)

	for i, row := range checkMatrixParallel {
		for j, val := range row {
			checkMatrixParallel[i][j] = math.Round(val*10000) / 10000
		}
	}

	err = fileProcessor.WriteCsv("../check_matrix_parallel.csv", matrix.StringifyMatrix(checkMatrixParallel))

	if err != nil {
		t.Fatal(err)
	}

}

func TestTempParallel(t *testing.T) {
	res := matrix.GenerateRandomMatrixInStringFormat(50)
	fileProcessor := file_processor.New()
	err := fileProcessor.WriteCsv("../matrix.csv", res)

	if err != nil {
		t.Fatal(err)
	}

	m, err := matrix.MakeFromFile("../matrix.csv", fileProcessor)

	if err != nil {
		t.Fatal(err)
	}

	inverseCalculator := New()

	inverse, err, ms := inverseCalculator.CalculateInverseMatrixParallel(m)

	fmt.Println(ms)

	err = fileProcessor.WriteCsv("../inverse_matrix.csv", matrix.StringifyMatrix(inverse))

	if err != nil {
		t.Fatal(err)
	}

	checkMatrix := MatrixMultiply(m.GetMatrix(), inverse)

	for i, row := range checkMatrix {
		for j, val := range row {
			checkMatrix[i][j] = math.Round(val*10000) / 10000
		}
	}

	err = fileProcessor.WriteCsv("../check_matrix.csv", matrix.StringifyMatrix(checkMatrix))

	if err != nil {
		t.Fatal(err)
	}

}

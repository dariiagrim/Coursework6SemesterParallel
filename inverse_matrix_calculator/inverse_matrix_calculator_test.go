package inverse_matrix_calculator

import (
	pkgMatrix "CourseWorkParallel/matrix"
	"os"
	"testing"
)

var calculator InverseMatrixCalculator

func TestMain(m *testing.M) {
	calculator = New()
	os.Exit(m.Run())
}

func TestCalculateInverseMatrixSequential(t *testing.T) {
	matrix := pkgMatrix.New([][]float64{
		{1, 2},
		{3, 4},
	})

	_, err := calculator.CalculateInverseMatrixSequential(matrix)
	if err != nil {
		t.Fatal(err)
	}
}

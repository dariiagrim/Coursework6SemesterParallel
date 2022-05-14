package inverse_matrix_calculator

/*
#cgo LDFLAGS: ./libmatrix_mul.a -ldl
#include "./libmatrix_mul.h"
*/
import "C"
import (
	pkgMatrix "CourseWorkParallel/matrix"
	"errors"
	"fmt"
	"unsafe"
)

type InverseMatrixCalculator struct{}

func New() InverseMatrixCalculator {
	return InverseMatrixCalculator{}
}

func (c *InverseMatrixCalculator) CalculateInverseMatrixSequential(matrix *pkgMatrix.Matrix) ([][]int64, error) {
	if len(matrix.GetMatrix()) == 0 || len(matrix.GetMatrix()) != len(matrix.GetMatrix()[0]) {
		return nil, errors.New("matrix is not square")
	}
	matrix.MakeExtended()
	for i := 0; i < len(matrix.GetMatrix()); i++ {
		num, err := matrix.FindFistNonZeroColumnRow(i, i, len(matrix.GetMatrix()))
		if err != nil {
			return nil, err
		}
		if num != i {
			matrix.SwapRows(num, i)
		}
		matrix.SetToZeroColumnUpToDown(i)
	}

	for i := len(matrix.GetMatrix()) - 1; i >= 0; i-- {
		if matrix.GetMatrix()[i][i] == 0 {
			num, err := matrix.FindFistNonZeroColumnRow(i, 0, i+1)
			if err != nil {
				return nil, err
			}
			if num != i {
				matrix.SwapRows(num, i)
			}
		}
		matrix.SetToZeroColumnDownToUp(i)
	}

	for i := 0; i < len(matrix.GetMatrix()); i++ {
		a := 1 / matrix.GetMatrix()[i][i]
		matrix.MultiplyRowByNumber(i, a)
	}

	oldMatrix := pkgMatrix.New([][]float64{
		{1, 2},
		{3, 4},
	})

	newData := make([][]float64, len(matrix.GetMatrix()))
	for i := 0; i < len(matrix.GetMatrix()); i++ {
		newData[i] = make([]float64, len(matrix.GetMatrix()))
		for j := 0; j < len(matrix.GetMatrix()); j++ {
			newData[i][j] = matrix.GetMatrix()[i][j+len(matrix.GetMatrix())]
		}
	}

	result := MatrixMultiply(oldMatrix.GetMatrix(), newData)

	for _, row := range result {
		fmt.Println(row)
	}

	return nil, nil
}

func MatrixMultiply(first, second [][]float64) [][]float64 {
	size := len(first)
	firstData := make([]float64, 0)
	secondData := make([]float64, 0)

	for _, row := range first {
		firstData = append(firstData, row...)
	}

	for _, row := range second {
		secondData = append(secondData, row...)
	}

	firstMatrix := C.matrix_new(C.ulong(size), C.ulong(size), (*C.double)(&firstData[0]))
	secondMatrix := C.matrix_new(C.ulong(size), C.ulong(size), (*C.double)(&secondData[0]))

	resultMatrix := C.multiply_async(firstMatrix, secondMatrix)

	resultData := make([]float64, resultMatrix.rows*resultMatrix.cols)

	pointerSize := unsafe.Sizeof(resultMatrix.data)

	for i := 0; i < len(resultData); i++ {
		pointer := (*C.double)(unsafe.Pointer(uintptr(unsafe.Pointer(resultMatrix.data)) + uintptr(i)*pointerSize))
		resultData[i] = (float64)(*pointer)
	}

	result := make([][]float64, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			result[i][j] = resultData[j+i*size]
		}
	}

	return result
}

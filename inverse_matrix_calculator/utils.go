package inverse_matrix_calculator

/*
#cgo LDFLAGS: ./libmatrix_mul.a -ldl
#include "./libmatrix_mul.h"
*/
import "C"
import "unsafe"

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

package inverse_matrix_calculator

import (
	pkgMatrix "CourseWorkParallel/matrix"
	"errors"
	"time"
)

type InverseMatrixCalculator struct{}

func New() InverseMatrixCalculator {
	return InverseMatrixCalculator{}
}

func (c *InverseMatrixCalculator) CalculateInverseMatrixSequential(matrix *pkgMatrix.Matrix) ([][]float64, error, int64) {
	startTime := time.Now()
	if len(matrix.GetMatrix()) == 0 || len(matrix.GetMatrix()) != len(matrix.GetMatrix()[0]) {
		return nil, errors.New("matrix is not square"), 0
	}
	matrix.MakeExtended()
	for i := 0; i < matrix.GetInitialSize(); i++ {
		num, err := matrix.FindFistNonZeroColumnRow(i, i, matrix.GetInitialSize())
		if err != nil {
			return nil, err, 0
		}
		if num != i {
			matrix.SwapRows(num, i)
		}
		matrix.SetToZeroColumnUpToDown(i)
	}

	for i := matrix.GetInitialSize() - 1; i >= 0; i-- {
		if matrix.GetExtendedMatrix()[i][i] == 0 {
			num, err := matrix.FindFistNonZeroColumnRow(i, 0, i+1)
			if err != nil {
				return nil, err, 0
			}
			if num != i {
				matrix.SwapRows(num, i)
			}
		}
		matrix.SetToZeroColumnDownToUp(i)
	}

	for i := 0; i < matrix.GetInitialSize(); i++ {
		a := 1 / matrix.GetExtendedMatrix()[i][i]
		matrix.MultiplyRowByNumber(i, a)
	}

	inverseMatrix := make([][]float64, 0)

	for _, val := range matrix.GetExtendedMatrix() {
		inverseMatrix = append(inverseMatrix, val[matrix.GetInitialSize():matrix.GetExtendedSize()])
	}

	return inverseMatrix, nil, time.Now().Sub(startTime).Microseconds()
}

func (c *InverseMatrixCalculator) CalculateInverseMatrixParallel(matrix *pkgMatrix.Matrix) ([][]float64, error, int64) {
	startTime := time.Now()
	if len(matrix.GetMatrix()) == 0 || len(matrix.GetMatrix()) != len(matrix.GetMatrix()[0]) {
		return nil, errors.New("matrix is not square"), 0
	}
	matrix.MakeExtended()
	for i := 0; i < matrix.GetInitialSize(); i++ {
		num, err := matrix.FindFistNonZeroColumnRow(i, i, matrix.GetInitialSize())
		if err != nil {
			return nil, err, 0
		}
		if num != i {
			matrix.SwapRows(num, i)
		}
		chans := matrix.SetToZeroColumnUpToDownParallel(i)
		for _, ch := range chans {
			_ = <-ch
		}
	}

	for i := matrix.GetInitialSize() - 1; i >= 0; i-- {
		if matrix.GetExtendedMatrix()[i][i] == 0 {
			num, err := matrix.FindFistNonZeroColumnRow(i, 0, i+1)
			if err != nil {
				return nil, err, 0
			}
			if num != i {
				matrix.SwapRows(num, i)
			}
		}
		chans := matrix.SetToZeroColumnDownToUpParallel(i)
		for _, ch := range chans {
			_ = <-ch
		}

	}

	chans := make([]chan bool, 0)
	for i := 0; i < matrix.GetInitialSize(); i++ {
		a := 1 / matrix.GetExtendedMatrix()[i][i]
		iCopy := i
		ch := make(chan bool, 2)
		chans = append(chans, ch)
		go func() {
			matrix.MultiplyRowByNumber(iCopy, a)
			ch <- true
		}()
	}

	for _, ch := range chans {
		_ = <-ch
	}

	inverseMatrix := make([][]float64, 0)

	for _, val := range matrix.GetExtendedMatrix() {
		inverseMatrix = append(inverseMatrix, val[matrix.GetInitialSize():matrix.GetExtendedSize()])
	}

	return inverseMatrix, nil, time.Now().Sub(startTime).Microseconds()
}

package matrix

import "errors"

type Matrix struct {
	matrix [][]float64
}

func New(matrix [][]float64) *Matrix {
	return &Matrix{
		matrix: matrix,
	}
}

func newIdentity(size int) *Matrix {
	matrix := make([][]float64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]float64, size)
		matrix[i][i] = 1
	}

	return &Matrix{
		matrix: matrix,
	}
}

func (m *Matrix) GetMatrix() [][]float64 {
	return m.matrix
}

func (m *Matrix) SwapRows(indexI int, indexJ int) {
	rowCopy := m.matrix[indexI]
	m.matrix[indexI] = m.matrix[indexJ]
	m.matrix[indexJ] = rowCopy
}

func (m *Matrix) FindFistNonZeroColumnRow(indexColumn int, fromRow int, toRow int) (int, error) {
	for i := fromRow; i < toRow; i++ {
		if m.matrix[i][indexColumn] != 0 {
			return i, nil
		}
	}
	return 0, errors.New("zero determinant")
}

func (m *Matrix) SetToZeroColumnUpToDown(indexColumn int) {
	a := m.matrix[indexColumn][indexColumn]
	for i := indexColumn + 1; i < len(m.matrix); i++ {
		b := m.matrix[i][indexColumn]
		if b == 0 {
			continue
		}
		coef := -a / b
		for j := 0; j < len(m.matrix[0]); j++ {
			m.matrix[i][j] = m.matrix[i][j]*coef + m.matrix[indexColumn][j]
		}
	}
}

func (m *Matrix) SetToZeroColumnDownToUp(indexColumn int) {
	a := m.matrix[indexColumn][indexColumn]
	for i := indexColumn - 1; i >= 0; i-- {
		b := m.matrix[i][indexColumn]
		if b == 0 {
			continue
		}
		coef := -a / b
		for j := 0; j < len(m.matrix[0]); j++ {
			m.matrix[i][j] = m.matrix[i][j]*coef + m.matrix[indexColumn][j]
		}
	}
}

func (m *Matrix) MultiplyRowByNumber(indexRow int, number float64) {
	for j := 0; j < len(m.matrix[0]); j++ {
		m.matrix[indexRow][j] *= number
	}
}

func (m *Matrix) MakeExtended() {
	identityMatrix := newIdentity(len(m.matrix))

	for i, _ := range m.matrix {
		m.matrix[i] = append(m.matrix[i], identityMatrix.matrix[i]...)
	}
}

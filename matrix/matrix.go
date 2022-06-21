package matrix

import (
	"CourseWorkParallel/file_processor"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type Matrix struct {
	initialMatrix  [][]float64
	extendedMatrix [][]float64
	initialSize    int
	extendedSize   int
}

func New(matrix [][]float64) *Matrix {
	return &Matrix{
		initialMatrix: matrix,
		initialSize:   len(matrix[0]),
		extendedSize:  len(matrix[0]) * 2,
	}
}

func StringifyMatrix(m [][]float64) [][]string {
	res := make([][]string, len(m))
	for i, row := range m {
		for _, val := range row {
			res[i] = append(res[i], fmt.Sprint(val))
		}
	}

	return res
}

func MakeFromFile(path string, fileProcessor *file_processor.FileProcessor) (*Matrix, error) {
	data, err := fileProcessor.ReadCsv(path)

	if err != nil {
		return nil, err
	}

	matrixArr := make([][]float64, len(data))

	for i, row := range data {
		for _, val := range row {
			number, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, err
			}
			matrixArr[i] = append(matrixArr[i], number)
		}
	}

	return New(matrixArr), nil

}

func GenerateRandomMatrixInStringFormat(size int) [][]string {
	rand.Seed(time.Now().UnixNano())
	res := make([][]string, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			res[i] = append(res[i], fmt.Sprint(math.Round((1+rand.Float64()*(99))*100)/100))
		}
	}

	return res
}

func NewIdentity(size int) *Matrix {
	matrix := make([][]float64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]float64, size)
		matrix[i][i] = 1
	}

	return &Matrix{
		initialMatrix: matrix,
	}
}

func (m *Matrix) GetMatrix() [][]float64 {
	return m.initialMatrix
}

func (m *Matrix) GetExtendedMatrix() [][]float64 {
	return m.extendedMatrix
}

func (m *Matrix) GetInitialSize() int {
	return m.initialSize
}

func (m *Matrix) GetExtendedSize() int {
	return m.extendedSize
}

func (m *Matrix) SwapRows(indexI int, indexJ int) {
	rowCopy := m.extendedMatrix[indexI]
	m.extendedMatrix[indexI] = m.extendedMatrix[indexJ]
	m.extendedMatrix[indexJ] = rowCopy
}

func (m *Matrix) FindFistNonZeroColumnRow(indexColumn int, fromRow int, toRow int) (int, error) {
	for i := fromRow; i < toRow; i++ {
		if m.extendedMatrix[i][indexColumn] != 0 {
			return i, nil
		}
	}
	return 0, errors.New("zero determinant")
}

func (m *Matrix) processRow(indexColumn int, indexRow int, a float64) {
	b := m.extendedMatrix[indexRow][indexColumn]
	if b == 0 {
		return
	}
	coef := -a / b
	for j := 0; j < len(m.extendedMatrix[0]); j++ {
		mul := m.extendedMatrix[indexRow][j] * coef
		m.extendedMatrix[indexRow][j] = mul + m.extendedMatrix[indexColumn][j]
	}
}

func (m *Matrix) SetToZeroColumnUpToDown(indexColumn int) {
	a := m.extendedMatrix[indexColumn][indexColumn]
	for i := indexColumn + 1; i < len(m.extendedMatrix); i++ {
		m.processRow(indexColumn, i, a)
	}
}

func (m *Matrix) SetToZeroColumnDownToUp(indexColumn int) {
	a := m.extendedMatrix[indexColumn][indexColumn]
	for i := indexColumn - 1; i >= 0; i-- {
		m.processRow(indexColumn, i, a)
	}
}

func (m *Matrix) MultiplyRowByNumber(indexRow int, number float64) {
	for j := 0; j < len(m.extendedMatrix[0]); j++ {
		m.extendedMatrix[indexRow][j] *= number
	}
}

func (m *Matrix) MakeExtended() {
	identityMatrix := NewIdentity(len(m.initialMatrix))
	m.extendedMatrix = make([][]float64, m.initialSize)

	for i, val := range m.initialMatrix {
		m.extendedMatrix[i] = append(m.extendedMatrix[i], val...)
		m.extendedMatrix[i] = append(m.extendedMatrix[i], identityMatrix.initialMatrix[i]...)
	}
}

func (m *Matrix) SetToZeroColumnUpToDownParallel(indexColumn int) []chan bool {
	chans := make([]chan bool, 0)
	a := m.extendedMatrix[indexColumn][indexColumn]
	for i := indexColumn + 1; i < len(m.extendedMatrix); i++ {
		iCopy := i
		ch := make(chan bool, 2)
		chans = append(chans, ch)
		go func() {
			defer close(ch)
			m.processRow(indexColumn, iCopy, a)
			ch <- true
		}()
	}
	return chans
}

func (m *Matrix) SetToZeroColumnDownToUpParallel(indexColumn int) []chan bool {
	chans := make([]chan bool, 0)
	a := m.extendedMatrix[indexColumn][indexColumn]
	for i := indexColumn - 1; i >= 0; i-- {
		iCopy := i
		ch := make(chan bool, 2)
		chans = append(chans, ch)
		go func() {
			defer close(ch)
			m.processRow(indexColumn, iCopy, a)
			ch <- true
		}()
	}
	return chans
}

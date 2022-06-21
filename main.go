package main

import (
	"CourseWorkParallel/file_processor"
	"CourseWorkParallel/inverse_matrix_calculator"
	"CourseWorkParallel/matrix"
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Custom?: yes -yes, else - no")
	var customString string
	custom := false
	_, _ = fmt.Scanln(&customString)
	if customString == "yes" {
		custom = true
	}
	fmt.Println("Select method: p - parallel, s - sequential, else - both")
	var method string
	_, err := fmt.Scanln(&method)
	if err != nil {
		method = "else"
	}
	size := 0
	fileProcessor := file_processor.New()
	if !custom {
		var sizeString string
		fmt.Println("Enter size: default 10")
		_, err = fmt.Scanln(&sizeString)
		if err != nil {
			sizeString = "10"
		}
		size, err = strconv.Atoi(sizeString)
		if err != nil {
			size = 10
		}
		res := matrix.GenerateRandomMatrixInStringFormat(size)
		if err := fileProcessor.WriteCsv("./matrix.csv", res); err != nil {
			panic(err)
		}
	}
	m, err := matrix.MakeFromFile("./matrix.csv", fileProcessor)
	if err != nil {
		panic(err)
	}
	inverseCalculator := inverse_matrix_calculator.New()

	switch method {
	case "s":
		doSeq(m, inverseCalculator, fileProcessor)
	case "p":
		doParallel(m, inverseCalculator, fileProcessor)
	default:
		doSeq(m, inverseCalculator, fileProcessor)
		doParallel(m, inverseCalculator, fileProcessor)
	}
}

func doSeq(m *matrix.Matrix, inverseCalculator inverse_matrix_calculator.InverseMatrixCalculator, fileProcessor *file_processor.FileProcessor) {
	inverse, err, ms := inverseCalculator.CalculateInverseMatrixSequential(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(float64(ms)/1000, "seq")
	if err := fileProcessor.WriteCsv("./inverse_matrix.csv", matrix.StringifyMatrix(inverse)); err != nil {
		panic(err)
	}
}

func doParallel(m *matrix.Matrix, inverseCalculator inverse_matrix_calculator.InverseMatrixCalculator, fileProcessor *file_processor.FileProcessor) {
	inverseParallel, err, ms := inverseCalculator.CalculateInverseMatrixParallel(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(float64(ms)/1000, "parallel")
	if err := fileProcessor.WriteCsv("./inverse_matrix_parallel.csv", matrix.StringifyMatrix(inverseParallel)); err != nil {
		panic(err)
	}
}

func measureTimeSeq() {
	fileProcessor := file_processor.New()
	sizes := []int{10, 20, 50, 100, 1000, 2000}
	for _, size := range sizes {
		times := make([]float64, 0)
		for i := 0; i < 5; i++ {
			res := matrix.GenerateRandomMatrixInStringFormat(size)
			if err := fileProcessor.WriteCsv("./matrix.csv", res); err != nil {
				panic(err)
			}
			m, err := matrix.MakeFromFile("./matrix.csv", fileProcessor)
			if err != nil {
				panic(err)
			}
			inverseCalculator := inverse_matrix_calculator.New()
			_, err, ms := inverseCalculator.CalculateInverseMatrixSequential(m)
			if err != nil {
				panic(err)
			}
			times = append(times, float64(ms)/1000)
		}
		avg := getAverage(times)
		fmt.Printf("For size %v, average time in 5 iterations %vms", size, avg)
		fmt.Println()
	}
}

func getAverage(arr []float64) float64 {
	sum := 0.0
	for _, el := range arr {
		sum += el
	}
	return math.Round(sum/float64(len(arr))*10000) / 10000
}

package file_processor

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFileProcessor_ReadCsv(t *testing.T) {
	fileProcessor := FileProcessor{}
	data, err := fileProcessor.ReadCsv("./test_file.csv")

	if err != nil {
		t.Fatal(err)
	}

	expectedData := [][]string{
		{"1", "2", "3.2", "4", "5"},
		{"5.5", "4", "3", "2", "1"},
	}

	a := assert.New(t)

	for i, row := range data {
		for j, val := range row {
			a.Equal(expectedData[i][j], val)
		}
	}
}

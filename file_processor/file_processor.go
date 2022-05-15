package file_processor

import (
	"encoding/csv"
	"os"
)

type FileProcessor struct{}

func New() *FileProcessor {
	return &FileProcessor{}
}

func (r *FileProcessor) ReadCsv(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (r *FileProcessor) WriteCsv(path string, records [][]string) error {
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			return err
		}
	}

	return nil
}

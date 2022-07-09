package reader

import (
	"encoding/csv"
	"io"
)

// Read file from given reader
func ReadFile(file io.Reader) ([]string, error) {
	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	var urls []string
	for _, line := range csvLines {
		urls = append(urls, line[0])
	}

	return urls, nil
}


// Package data provides utilities for loading the supplied detection data.
//
// The SE2/SE3/SE4 CSV files provided by Eye Security live in this directory.
// ReadCSV deliberately returns untyped records: how the detection data is
// modelled is left entirely to you.
package data

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ReadCSV reads the CSV file at path and returns one record per data row,
// each keyed by the column names in the header row. It imposes no schema on
// the data. An empty file yields an empty (non-nil) slice.
func ReadCSV(path string) ([]map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open csv %q: %w", path, err)
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read csv %q: %w", path, err)
	}

	records := make([]map[string]string, 0)
	if len(rows) == 0 {
		return records, nil
	}

	header := rows[0]
	for _, row := range rows[1:] {
		record := make(map[string]string, len(header))
		for i, col := range header {
			if i < len(row) {
				record[col] = row[i]
			}
		}
		records = append(records, record)
	}
	return records, nil
}

package office

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/h2non/filetype"
	"github.com/xuri/excelize/v2"
)

// Read
// only support csv/xls/xlsx/xlsm/xltx/xltm
func Read(r io.Reader, fn func([]string)) error {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return err
	}
	r = io.NopCloser(&buf)

	if filetype.IsDocument(buf.Bytes()) {
		return ReadExcel(r, fn)
	}

	return ReadCSV(r, fn)
}

func ReadExcel(r io.Reader, fn func([]string)) error {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}

	rows, err := f.Rows("Sheet1")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			continue
		}
		fn(row)
	}

	return nil
}

func ReadCSV(r io.Reader, fn func([]string)) error {
	cr := csv.NewReader(r)
	for {
		row, err := cr.Read()
		if err == io.EOF {
			break
		}
		fn(row)
	}
	return nil
}

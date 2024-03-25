package office

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"

	"github.com/h2non/filetype"
	"github.com/xuri/excelize/v2"
)

// Read
// only support csv/xls/xlsx/xlsm/xltx/xltm
func ReadExcel(r io.Reader, fn func([]string, error)) error {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return err
	}
	r = io.NopCloser(&buf)

	if filetype.IsDocument(buf.Bytes()) {
		return ReadXLS(r, fn)
	}

	ReadCSV(r, fn)
	return nil
}

func ReadXLS(r io.Reader, fn func([]string, error)) error {
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
		fn(rows.Columns())
	}

	return nil
}

func ReadCSV(r io.Reader, fn func([]string, error)) {
	cr := csv.NewReader(r)
	for {
		row, err := cr.Read()
		if errors.Is(err, io.EOF) {
			fn(nil, nil)
			return
		}
		fn(row, err)
	}
}

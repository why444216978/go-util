package extractor

import (
	"bytes"

	"github.com/ledongthuc/pdf"
)

func Extract(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}

	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func ExtractWithStyle(path string) (res []pdf.Text, err error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
	}()

	totalPage := r.NumPage()

	res = make([]pdf.Text, 0)
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		res = append(res, p.Content().Text...)
	}

	return
}

func ExtractRows(path string) (res []string, err error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
	}()

	totalPage := r.NumPage()

	res = make([]string, 0)
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, err := p.GetTextByRow()
		if err != nil {
			continue
		}
		l := len(rows)
		for i := l - 1; i >= 0; i-- {
			line := ""
			for _, word := range rows[i].Content {
				line += word.S
			}
			res = append(res, line)
		}
	}
	return
}

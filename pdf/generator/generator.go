package generator

import (
	"bytes"
	"os"
	"strings"

	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/pkg/errors"
)

// GeneratFromURL generat pdf from url
func GeneratFromURL(url, f string) error {
	file, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	pdfg, err := pdf.NewPDFGenerator()
	if err != nil {
		return errors.Wrap(err, "GeneratFromURL NewPDFGenerator err")
	}

	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(pdf.OrientationLandscape)
	pdfg.Grayscale.Set(true)

	page := pdf.NewPage(url)

	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)

	pdfg.SetOutput(file)
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return errors.Wrap(err, "GeneratFromURL Create err")
	}

	return nil

	err = pdfg.WriteFile(f)
	if err != nil {
		return errors.Wrap(err, "GeneratFromURL WriteFile err")
	}

	return nil
}

// GeneratFromHTML generat pdf from html
func GeneratFromHTML(html, f string) error {
	file, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	pdfg := pdf.NewPDFPreparer()
	pdfg.AddPage(pdf.NewPageReader(strings.NewReader(html)))
	pdfg.Dpi.Set(600)

	jsonBytes, err := pdfg.ToJSON()
	if err != nil {
		return errors.Wrap(err, "GeneratFromHTML ToJSON err")
	}

	pdfgFromJSON, err := pdf.NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		return errors.Wrap(err, "GeneratFromHTML NewPDFGeneratorFromJSON err")
	}

	pdfgFromJSON.SetOutput(file)
	// Create the PDF
	err = pdfgFromJSON.Create()
	if err != nil {
		return errors.Wrap(err, "GeneratFromHTML Create err")
	}
	return nil

	err = pdfgFromJSON.WriteFile(f)
	if err != nil {
		return errors.Wrap(err, "GeneratFromHTML WriteFile err")
	}

	return nil
}

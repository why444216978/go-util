package zip

import (
	"io"
	"io/fs"
	"os"

	"github.com/alexmullins/zip"
	"github.com/pkg/errors"
)

type File struct {
	Name   string
	Reader io.Reader
	Info   fs.FileInfo
}

func CreateEncryptedZip(files []File, output, password string) error {
	fzip, err := os.Create(output)
	if err != nil {
		return err
	}
	defer fzip.Close()

	for i, f := range files {
		if err := appendFile(fzip, f, password); err != nil {
			return errors.Wrapf(err, "append %d-%s fail", i, f.Name)
		}
	}

	return nil
}

func appendFile(fzip *os.File, f File, password string) error {
	header, err := zip.FileInfoHeader(f.Info)
	if err != nil {
		return errors.Wrap(err, "FileInfoHeader fail")
	}
	header.Name = f.Name
	header.Method = zip.Deflate

	zw := zip.NewWriter(fzip)
	defer zw.Close()

	w, err := zw.Encrypt(header.Name, password)
	if err != nil {
		return errors.Wrap(err, "Encrypt fail")
	}

	if _, err = io.Copy(w, f.Reader); err != nil {
		return errors.Wrap(err, "Copy fail")
	}
	return zw.Flush()
}

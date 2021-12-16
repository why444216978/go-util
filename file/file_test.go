package file

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

func TestWriteWithIo(t *testing.T) {
}

func TestReadLimit(t *testing.T) {
	Convey("TestReadLimit", t, func() {
		Convey("success has chinese", func() {
			r := strings.NewReader("啊1234567890abcde")
			res := ReadLimit(r, 10)
			assert.Equal(t, len(res), 10)
			assert.Equal(t, res, "啊1234567")
		})
		Convey("success no chinese", func() {
			r := strings.NewReader("1234567890abcde")
			res := ReadLimit(r, 10)
			assert.Equal(t, len(res), 10)
			assert.Equal(t, res, "1234567890")
		})
	})
}

func TestReadFile(t *testing.T) {

}

func TestReadFileLine(t *testing.T) {

}

func TestReadFromReader(t *testing.T) {

}

func TestReadJsonFile(t *testing.T) {

}

func TestGetFileInfo(t *testing.T) {

}

func TestGetFileMode(t *testing.T) {

}

func TestGetFileStat(t *testing.T) {

}

func TestChown(t *testing.T) {

}

func TestChmod(t *testing.T) {

}

func TestOpen(t *testing.T) {

}

func TestCreate(t *testing.T) {

}

func TestCleanFile(t *testing.T) {

}

func TestDownloadFileToBase64(t *testing.T) {

}

func TestBase64ToFile(t *testing.T) {

}

func TestMultiWriter(t *testing.T) {

}

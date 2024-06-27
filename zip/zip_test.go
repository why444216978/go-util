package zip

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEncryptedZip(t *testing.T) {
	f, _ := os.Open("./f.txt")
	defer f.Close()

	info, _ := f.Stat()

	err := CreateEncryptedZip([]File{
		{Name: "f.txt", Reader: f, Info: info},
	}, "data.zip", "123")
	assert.Nil(t, err)
}

package rsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	pubKey = `
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAz+NMbmcGyJyuN0tJFwBj
VXlQ61mpzOWiaGH05MTil0LeEecXy1ZoaKiND1cQwqiIFDgyOcc3XUxdSprGA29x
YGYEHud+HM8ZCeGO8LIWke4tiTOvpBlkiXQPQ9RUs+P6mFfq9eUzONKIOD+6GJXC
w8+hEJBIQeRIPfwq4VD8NZahruJRsaHFZLg86c5J469y4C9MODeYEDsmLElx3VW4
lXXABPUgTw8A7IxHYsWwzODH7Y/QQA4Cr1MtFyQkgz/LT98unW9qk+mAGh2DHgdl
fQ+RpVmZ+toTXEABF4hYNyecPd+FCNc53vqE1IcASrw4QwCOn9ihP9lYr7mdmacp
9ukMIuv2IML0nW/joZInnWIPexuu9zbwvkJ55cBCOc3LmSI9S8ePJH7XWTRrcs3E
m1uu5knudNy1Q02a0ztKeq40UKErQCifajZHy/K7ZxNx/ewm0Hhft4gmgHLx/Ptr
CCWTv0Te2rb5P30msinu48RQcSNDB58CfJ58C8Us/soNzwadJB3P8WcEW15Nm8h3
zKjz7ojAh+fEspm6ZW2nqXYYuLNbM3i+IfpZSYX5DNTXwNahTv4wmMxo98WRRFJ5
HIm5WT/UIZKM0EOBOBnByVj36fUDYF0r7FOka/1O4+vQTeZnkm6kid+W9ZvDlf6a
9oJPRC/SfTrw4cP1Kpo5b2UCAwEAAQ==
-----END PUBLIC KEY-----`
	privKey = `
-----BEGIN PRIVATE KEY-----
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
kErXhXDwV81bSxfp1sVTOBdxtp25WQ==
-----END PRIVATE KEY-----`
)

func TestEncrypt(t *testing.T) {
	data := `{"data":1}`
	res, err := PublicEncrypt(data, []byte(pubKey))
	assert.Nil(t, err)

	res, err = PrivateDecrypt(res, []byte(privKey))
	assert.Nil(t, err)
	assert.Equal(t, data, res)
}

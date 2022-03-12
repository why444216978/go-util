package hash

import (
	"hash/fnv"
)

// StringToUint32 string hash to uint32
func StringToUint32(str string) (uint32, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(str))
	if err != nil {
		return 0, err
	}

	return (h.Sum32() & 0x7FFFFFFF), nil
}

// StringToUint64 string hash to uint64
func StringToUint64(str string) (uint64, error) {
	h := fnv.New64a()
	_, err := h.Write([]byte(str))
	if err != nil {
		return 0, err
	}

	return (h.Sum64() & 0x7FFFFFFF), nil
}

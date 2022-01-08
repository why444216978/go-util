package snowflake

import (
	"hash/fnv"

	"github.com/bwmarrin/snowflake"

	"github.com/why444216978/go-util/sys"
)

var generator *snowflake.Node

func init() {
	var err error
	defer func() {
		if err == nil {
			return
		}
		generator, _ = snowflake.NewNode(1)
	}()

	n, err := getNode()
	if err != nil {
		return
	}

	generator, err = snowflake.NewNode(int64(n) % 1024)
	if err != nil {
		return
	}
}

func Generate() snowflake.ID {
	return generator.Generate()
}

func getNode() (uint32, error) {
	ip, err := sys.LocalIP()
	if err != nil {
		return 0, err
	}
	h := fnv.New32a()
	_, err = h.Write([]byte(ip))
	if err != nil {
		return 0, err
	}
	return (h.Sum32() & 0x7FFFFFFF), nil
}

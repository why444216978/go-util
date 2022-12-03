package net

import (
	"net"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func ipTrue(i net.IP) bool  { return true }
func ipFalse(i net.IP) bool { return false }

func TestIsInternalIP(t *testing.T) {
	convey.Convey("TestIsInternalIP", t, func() {
		assert.Equal(t, false, IsInternalIP(net.ParseIP("192.170.223.1"), nil, nil))
		assert.Equal(t, true, IsInternalIP(net.ParseIP("192.170.223.1"), ipTrue, ipTrue))
		assert.Equal(t, false, IsInternalIP(net.ParseIP("192.170.223.1"), ipFalse, ipFalse))
		assert.Equal(t, true, IsInternalIP(net.ParseIP("1030::C9B4:FF12:48AA:1A2B"), ipTrue, ipTrue))
		assert.Equal(t, false, IsInternalIP(net.ParseIP("1030::C9B4:FF12:48AA:1A2B"), ipFalse, ipFalse))
	})
}

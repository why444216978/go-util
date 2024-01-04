package net

import (
	"net"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestIsInternalIP(t *testing.T) {
	convey.Convey("TestIsInternalIP", t, func() {
		assert.Equal(t, false, IsInternalIP(net.ParseIP("192.170.223.1")))
		assert.Equal(t, true, IsInternalIP(net.ParseIP("172.22.152.242")))
		assert.Equal(t, false, IsInternalIP(net.ParseIP("1030::C9B4:FF12:48AA:1A2B")))
	})
}

package iptree

import (
	"bytes"
	"testing"
)

// var v4InV6Prefix = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}

// func Test_parse_invalid(t *testing.T) {
// 	a, b, err := Parse("hello world")
// 	if err == nil {
// 		t.Error("parse invalid error")
// 	}
// 	if a != nil {
// 		t.Error("parse invalid error")
// 	}
// 	if b != nil {
// 		t.Error("parse invalid error")
// 	}
// }

// func Test_parse_ipv4(t *testing.T) {
// 	a, b, err := Parse("1.1.1.1")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if b != nil {
// 		t.Error("ip v4 error")
// 	}
// 	if a == nil {
// 		t.Error("ip v4 error")
// 		return
// 	}

// 	if len(*a) != 16 {
// 		t.Error("ip v4 error")
// 	}

// 	if !bytes.Equal(*a, []byte{
// 		0, 0, 0, 0,
// 		0, 0, 0, 0,
// 		0, 0, 0xff, 0xff,
// 		0x01, 0x01, 0x01, 0x01}) {
// 		t.Error("ip v4 error")
// 	}
// }

// func Test_parse_ipv6(t *testing.T) {
// 	a, b, err := Parse("::1")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if b != nil {
// 		t.Error("ip v6 error")
// 	}
// 	if a == nil {
// 		t.Error("ip v6 error")
// 		return
// 	}

// 	if !bytes.Equal(*a, []byte{
// 		0, 0, 0, 0,
// 		0, 0, 0, 0,
// 		0, 0, 0, 0,
// 		0, 0, 0, 0x01}) {
// 		t.Error("ip v6 error")
// 	}
// }

// func Test_parse_ipv4_cidr(t *testing.T) {
// 	a, b, err := Parse("192.168.0.1/11")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if b == nil {
// 		t.Error("ip v4 cidr error")
// 	}
// 	if a != nil {
// 		t.Error("ip v4 cidr error")
// 		return
// 	}

// 	if !bytes.Equal(b[0], []byte{
// 		0, 0, 0, 0,
// 		0, 0, 0, 0,
// 		0, 0, 0xff, 0xff,
// 		192, 160, 0, 0,
// 	}) {
// 		t.Error("ip v4 cidr error")
// 	}

// 	if !bytes.Equal(b[1], []byte{
// 		0, 0, 0, 0,
// 		0, 0, 0, 0,
// 		0, 0, 0xff, 0xff,
// 		192, 191, 0xff, 0xff,
// 	}) {
// 		t.Error("ip v4 cidr error")
// 	}
// }

// func Test_parse_ipv6_cidr(t *testing.T) {
// 	a, b, err := Parse("2001:0DB8:0000:0023:0008:0800:200C:417A/67")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if b == nil {
// 		t.Error("ip v6 prefix error")
// 	}
// 	if a != nil {
// 		t.Error("ip v6 prefix error")
// 		return
// 	}

// 	if !bytes.Equal(b[0], []byte{
// 		0x20, 0x01, 0x0d, 0xb8,
// 		0, 0, 0, 0x23,
// 		0, 0, 0, 0,
// 		0, 0, 0, 0,
// 	}) {
// 		t.Error("ip v6 prefix error")
// 	}

// 	if !bytes.Equal(b[1], []byte{
// 		0x20, 0x01, 0x0d, 0xb8,
// 		0, 0, 0, 0x23,
// 		0x1f, 0xff, 0xff, 0xff,
// 		0xff, 0xff, 0xff, 0xff,
// 	}) {
// 		t.Error("ip v6 prefix error")
// 	}
// }

func Test_parse_ipv4_range(t *testing.T) {
	a, b, err := Parse("1.2.3.4-4.3.2.1")
	if err != nil {
		t.Error(err)
	}
	if b == nil {
		t.Error("ip v4 range error")
	}
	if a != nil {
		t.Error("ip v4 range error")
		return
	}

	if !bytes.Equal(b[0], []byte{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0xff, 0xff,
		1, 2, 3, 4,
	}) {
		t.Error("ip v4 range error")
	}

	if !bytes.Equal(b[1], []byte{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0xff, 0xff,
		4, 3, 2, 1,
	}) {
		t.Error("ip v4 range error")
	}
}

func Test_parse_ipv6_range(t *testing.T) {
	a, b, err := Parse("::1-2::")
	if err != nil {
		t.Error(err)
	}
	if b == nil {
		t.Error("ip v6 range error")
	}
	if a != nil {
		t.Error("ip v6 range error")
		return
	}

	if !bytes.Equal(b[0], []byte{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 1,
	}) {
		t.Error("ip v6 range error")
	}

	if !bytes.Equal(b[1], []byte{
		0, 2, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	}) {
		t.Error("ip v6 range error")
	}
}

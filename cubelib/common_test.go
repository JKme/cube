package cubelib

import "testing"

func TestName(t *testing.T) {
	ParseService("127.0.0.1:22")
}

func TestParseNet(t *testing.T) {
	ParseNet("ftp://127.0.0.1:22")
}

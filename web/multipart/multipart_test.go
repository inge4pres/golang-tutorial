package mpart

import (
	"os"
	"testing"
)

var boundary = "--.--MyBoundary--.--"
var text = []byte("This MN has no sense\r\n")
var machine = []byte("Reporting-UA: EASYAS2;\r\nMDN-Status: processed;\r\n")

func TestWriteUnsignedMdnReport(t *testing.T) {
	err := WriteUnsignedMdnReport(os.Stdout, boundary, text, machine)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestWriteSignedMdnReport(t *testing.T) {
	err := WriteSignedMdnReport(os.Stdout, boundary, text, machine, []byte("ULTRGASIGNATUREbyteArray"))
	if err != nil {
		t.Errorf("%v", err)
	}
}

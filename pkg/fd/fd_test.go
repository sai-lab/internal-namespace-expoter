package fd

import (
	"bytes"
	"testing"
)

func TestLoadNumberOfFiles(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("1088\t0\t9223372036854775807")

	numberOfFiles, err := ParseNumberOfFiles(&buf)
	if err != nil {
		t.Error(err)
	}

	if numberOfFiles.usingFd != 1088 {
		t.Errorf("Assertion Error: expect(%d), actual(%d)", 1088, numberOfFiles.usingFd)
	}

	if numberOfFiles.availableFd != 0 {
		t.Errorf("Assertion Error: expect(%d), actual(%d)", 0, numberOfFiles.availableFd)
	}
	if numberOfFiles.maxFd != 9223372036854775807 {
		t.Errorf("Assertion Error: expect(%d), actual(%d)", 9223372036854775807, numberOfFiles.maxFd)
	}

}

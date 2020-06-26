package fd

import (
	"fmt"
	"io"
	"os"
)

type NumberOfFiles struct {
	usingFd     int64
	availableFd int64
	maxFd       int64
}

const fileNr = "/proc/sys/fs/file-nr"

func OpenFileNr() (io.Reader, error) {
	file, err := os.Open(fileNr)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ParseNumberOfFiles(r io.Reader) (*NumberOfFiles, error) {

	buf := make([]byte, 255)

	_, err := r.Read(buf)
	if err != nil {
		return nil, err
	}

	line := string(buf)
	fmt.Println(line)

	numberOfFiles := NumberOfFiles{}
	fmt.Sscanf(line, "%d\t%d\t%d", &numberOfFiles.usingFd, &numberOfFiles.availableFd, &numberOfFiles.maxFd)
	return &numberOfFiles, nil
}

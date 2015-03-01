package main

import "io"

type teeReader struct {
	reader io.Reader
	writer io.Writer
}

func (tee *teeReader) Read(portion []byte) (bytesRead int, err error) {
	bytesRead, err = tee.reader.Read(portion)
	if bytesRead > 0 {
		if bytesRead, err := tee.writer.Write(portion[:bytesRead]); err != nil {
			return bytesRead, err
		}
	}
	return
}

func main() {}

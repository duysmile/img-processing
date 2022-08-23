package convert

import (
	"github.com/h2non/bimg"
	"io"
)

func ConvertFromFile(inputFile string) (output []byte, err error) {
	buffer, err := bimg.Read(inputFile)
	if err != nil {
		return
	}

	output, err = bimg.NewImage(buffer).Convert(bimg.JPEG)
	return
}

func ConvertFromBuf(in io.Reader) (out io.Writer) {
	return
}

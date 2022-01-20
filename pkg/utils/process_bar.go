package utils

import (
	"github.com/schollz/progressbar/v3"
	"io"
)

func PrintWithProcessBar(bytesLength int64, description string, body io.ReadCloser, writers ...io.Writer) error {
	bar := progressbar.DefaultBytes(
		bytesLength,
		description,
	)

	if writers == nil {
		writers = make([]io.Writer, 0, 1)
	}
	writers = append(writers, bar)
	_, err := io.Copy(io.MultiWriter(writers...), body)
	if err != nil {
		return err
	}
	return nil
}

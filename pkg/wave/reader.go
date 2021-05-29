package wave

import (
	"io"
)

func ReadWave(input io.Reader, output chan []byte, bufsize int) error {
	defer close(output)
	for {
		data := make([]byte, bufsize)
		n, err := input.Read(data)
		if n > 0 {
			output <- data[0:n]
		}
		if err != nil {
			return err
		}
	}
}

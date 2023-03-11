package server

import "io"

func Decode(in []byte, out io.Reader) error {
	_, err := out.Read(in)
	if err != nil {
		return err
	}

	return nil
}

package container

import (
	"archive/tar"
	"fmt"
)

func CreateTarHeader(name string, mode int64, size int64) (*tar.Header, error) {
	return &tar.Header{
		Name: name,
		Mode: mode,
		Size: size,
	}, nil
}

func WriteTarFileContent(tw *tar.Writer, content []byte) error {
	if _, err := tw.Write(content); err != nil {
		return fmt.Errorf("error writing file content: %w", err)
	}
	return nil
}

func CloseTarWriter(tw *tar.Writer) error {
	if err := tw.Close(); err != nil {
		return fmt.Errorf("error closing tar writer: %w", err)
	}
	return nil
}

package transporter

import (
	"bufio"
	"fmt"
	"os"
)

// FileSource struct for handling file-based imports
type FileSource struct {
	filePath string
	file     *os.File
	reader   *bufio.Reader
}

// NewFileSource initializes and returns a new FileSource
func NewFileSource(filePath string) *FileSource {
	return &FileSource{
		filePath: filePath,
	}
}

// openFile opens the file for reading if it's not already opened.
func (f *FileSource) openFile() error {
	if f.file == nil {
		file, err := os.Open(f.filePath)
		if err != nil {
			return err
		}
		f.file = file
		f.reader = bufio.NewReader(file)
	}
	return nil
}

// ReadData reads a line of data from the file.
func (f *FileSource) ReadData() (string, error) {
	err := f.openFile()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	line, err := f.reader.ReadString('\n')
	if err != nil {
		return "", err // End of file or read error
	}

	return line, nil
}

// Close closes the file handle.
func (f *FileSource) Close() error {
	if f.file != nil {
		err := f.file.Close()
		f.file = nil
		f.reader = nil
		return err
	}
	return nil
}

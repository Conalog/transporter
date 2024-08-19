package transporter

import (
	"fmt"
	"os"
)

// FileDestination struct for handling file-based exports
type FileDestination struct {
	filePath              string
	appendNewLineEachData bool
	file                  *os.File
}

// NewFileDestination initializes and returns a new FileDestination
// The file is not opened immediately.
func NewFileDestination(filePath string, appendNewLineEachData bool) *FileDestination {
	return &FileDestination{
		filePath:              filePath,
		appendNewLineEachData: appendNewLineEachData,
	}
}

// openFile opens the file in append mode if it's not already opened.
func (f *FileDestination) openFile() error {
	if f.file == nil {
		// Open file in append mode, which allows other processes to read while writing.
		file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		f.file = file
	}
	return nil
}

// WriteData writes the buffered data to the file.
// It attempts to open the file if it's not already opened and skips writing if an error occurs.
func (f *FileDestination) WriteData(data string) error {
	err := f.openFile()
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return nil // Skip writing if unable to open file
	}

	_, err = f.file.WriteString(data)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		f.file.Close() // Close broken file handle
		f.file = nil   // Reset file handle for next attempt
	}

	// Append a new line after each data if enabled
	if f.appendNewLineEachData {
		_, err = f.file.WriteString("\n")
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			f.file.Close() // Close broken file handle
			f.file = nil   // Reset file handle for next attempt
		}
	}

	return nil
}

// Close closes the file handle.
func (f *FileDestination) Close() error {
	if f.file != nil {
		err := f.file.Close()
		f.file = nil
		return err
	}
	return nil
}

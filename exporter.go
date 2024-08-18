package transporter

import (
	"context"
	"fmt"
)

// Destination interface that all destinations must implement
type Destination interface {
	WriteData(data string) error
	Close() error
}

// DataExporter struct that manages exporting data to multiple destinations
type DataExporter struct {
	dataChan     chan string
	destinations []Destination
}

// NewDataExporter initializes and returns a new DataExporter
func NewDataExporter(dataChan chan string, destinations []Destination) *DataExporter {
	return &DataExporter{
		dataChan:     dataChan,
		destinations: destinations,
	}
}

// Start begins reading data from the channel and sends it to all destinations
func (de *DataExporter) Start(ctx context.Context) {
	defer de.Close()

	// Continuously read from the channel until it's closed
	for {
		select {
		case data, ok := <-de.dataChan:
			if !ok {
				return
			}
			de.exportData(data)
		case <-ctx.Done():
			return
		}
	}
}

// exportData sends the data to all destinations
func (de *DataExporter) exportData(data string) {
	for _, dest := range de.destinations {
		err := dest.WriteData(data)
		if err != nil {
			fmt.Println("Error writing data:", err)
		}
	}
}

// Close closes all destinations managed by the DataExporter
func (de *DataExporter) Close() {
	for _, dest := range de.destinations {
		err := dest.Close()
		if err != nil {
			fmt.Println("Error closing destination:", err)
		}
	}
}

/*
func main() {
	dataChan := make(chan string, 10)

	// Create file destination (file will be opened on first write attempt)
	fileDest := NewFileDestination("output.txt")
	defer fileDest.Close()

	// Create socket destination (connection will be established on first write attempt)
	socketDest := NewSocketDestination("localhost:8080")
	defer socketDest.Close()

	// Create DataExporter with both destinations
	exporter := NewDataExporter(dataChan, []Destination{fileDest, socketDest})
	go exporter.Start()

	// Example of sending data to the channel
	dataChan <- "Hello, World!\n"
	dataChan <- "More data...\n"

	// Close the channel when done
	close(dataChan)

	// Ensure all destinations are properly closed
	exporter.Close()
}
*/

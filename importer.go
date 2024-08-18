package transporter

import (
	"context"
	"fmt"
)

// Source interface that all sources must implement
type Source interface {
	ReadData() (string, error)
	Close() error
}

// DataImporter struct that manages importing data from multiple sources
type DataImporter struct {
	dataChan chan string
	sources  []Source
}

// NewDataImporter initializes and returns a new DataImporter
func NewDataImporter(dataChan chan string, sources []Source) *DataImporter {
	return &DataImporter{
		dataChan: dataChan,
		sources:  sources,
	}
}

// Start begins reading data from all sources and sends it to the data channel
func (di *DataImporter) Start(ctx context.Context) {
	for _, source := range di.sources {
		go di.readFromSource(ctx, source)
	}
}

// readFromSource continuously reads from a single source and sends data to the channel
func (di *DataImporter) readFromSource(ctx context.Context, source Source) {
	defer source.Close()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			data, err := source.ReadData()
			if err != nil {
				fmt.Println("Error reading data:", err)
				break
			}
			di.dataChan <- data
		}
	}
}

// Close closes all sources managed by the DataImporter
func (di *DataImporter) Close() {
	close(di.dataChan)
	for _, source := range di.sources {
		err := source.Close()
		if err != nil {
			fmt.Println("Error closing source:", err)
		}
	}
}

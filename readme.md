# Transporter

## Description

This is a simple transporter that can receive and send messages from various sources, with go channels.
Currently, string data is supported and newline is used as a delimiter. Try using hexadecimal strings for binary data.

## Supported sources

### Exporter

- [x] Socket
- [x] File
- [ ] HTTP

### Importer

- [x] Socket
- [ ] File
- [ ] HTTP

## Usage

```go
package main

func main(){
    // Intialize context for stopping the transporter
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // 1. Start the exporter
    rawDataExportChan := make(chan string, 8192)
    rawDataDestinations := []Destination{
        NewFileDestination("path/to/file"),
        NewSocketDestination("tcp://localhost:8080"),
    }
    rawDataExporter := NewDataExporter(rawDataExportChan, rawDataDestinations)
    go rawDataExporter.Start(ctx) // Close function is called in the defer statement

    // 2. Start the importer
    commandDataChan := make(chan string, 50000)
    commandSources := []Source{
        NewSocketSource("tcp://localhost:8080"),
    }
    commandImporter := NewDataImporter(commandDataChan, commandSources)
    go commandImporter.Start(ctx)

    // Do something with the data
    commandDataChan <- "0A 0B 0C 0D"
}
```

## Address format

- File: `/path/to/file`
- Socket(TCP): `tcp://localhost:8080`
- Socket(UDP): `udp://localhost:8080`
- Socket(Unix): `unix:///path/to/socket`, `/path/to/socket`

## License

MIT

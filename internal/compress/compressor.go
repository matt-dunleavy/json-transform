// internal/compress/compressor.go
package compress

import (
    "compress/gzip"
    "io"
    "os"
)

// CompressJSONFile compresses the given JSON file to a .gz file.
func CompressJSONFile(inputFile, outputFile string) error {
    input, err := os.Open(inputFile)
    if err != nil {
        return err
    }
    defer input.Close()

    output, err := os.Create(outputFile)
    if err != nil {
        return err
    }
    defer output.Close()

    writer := gzip.NewWriter(output)
    defer writer.Close()

    _, err = io.Copy(writer, input)
    if err != nil {
        return err
    }

    return nil
}

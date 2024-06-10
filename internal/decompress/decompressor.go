// internal/decompress/decompressor.go
package decompress

import (
    "compress/gzip"
    "io"
    "os"
)

// DecompressJSONFile decompresses the given .gz file to a JSON file.
func DecompressJSONFile(inputFile, outputFile string) error {
    input, err := os.Open(inputFile)
    if err != nil {
        return err
    }
    defer input.Close()

    reader, err := gzip.NewReader(input)
    if err != nil {
        return err
    }
    defer reader.Close()

    output, err := os.Create(outputFile)
    if err != nil {
        return err
    }
    defer output.Close()

    _, err = io.Copy(output, reader)
    if err != nil {
        return err
    }

    return nil
}

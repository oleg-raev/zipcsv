package zipcsv

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

const ZIP_EXT = ".zip"
const CSV_EXT = ".csv"

func ProcessFiles(files []string) (<-chan string, <-chan error) {
	row := make(chan string, 1)
	error := make(chan error)

	go func() {
		for _, zipFile := range files {
			reader, err := zip.OpenReader(zipFile)
			defer reader.Close()
			if err != nil {
				error <- err
			}

			for _, file := range reader.File {
				if !IsCSV(file.Name) {
					continue
				}

				fileReader, err := file.Open()
				defer fileReader.Close()
				if err != nil {
					if fileReader != nil {
						fileReader.Close()
					}
					error <- err
				}

				reader := bufio.NewReader(fileReader)
				for {
					line, _, err := reader.ReadLine()
					if err != nil {
						if err == io.EOF {
							fmt.Println("CLOSED")
							break
						}
						error <- err
					}

					row <- string(line)
				}

				fileReader.Close()
			}
			reader.Close()
		}
		close(row)
		close(error)
	}()

	return row, error
}

func IsZIP(name string) bool {
	return ZIP_EXT == strings.ToLower(filepath.Ext(name))
}

func IsCSV(name string) bool {
	return CSV_EXT == strings.ToLower(filepath.Ext(name))
}

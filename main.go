package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	result := make(map[string]int64)

	rows, errs := inputDataProcessing()

loop:
	for {
		select {
		case row, ok := <-rows:
			if !ok {
				break loop
			}
			processRow(row, result)
		case err, ok := <-errs:
			if !ok {
				break loop
			}
			fmt.Println(err)
		}
	}

	for key, val := range result {
		fmt.Printf("%s => %d\n", key, val)
	}
}

func inputDataProcessing() (<-chan string, <-chan error) {
	var (
		chOut = make(chan string)
		chErr = make(chan error)
	)

	in, err := gzip.NewReader(os.Stdin)
	if err != nil {
		panic(err)
	}

	go func() {
		for err := range chErr {
			//TODO handle errors
			fmt.Println(err)
		}
	}()

	go func() {
		reader := bufio.NewReader(in)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				if err == io.EOF {
					fmt.Println("CLOSED")
					break
				}
				chErr <- err
			}
			fmt.Println(string(line))
			chOut <- string(line)
		}
		close(chErr)
		close(chOut)
	}()

	return chOut, chErr
}

func processRow(row string, result map[string]int64) {
	data := strings.Split(row, ";")
	if len(data) < 1 || len(data[0]) < 13 {
		return
	}

	key := data[0][11:13]

	if _, ok := result[key]; !ok {
		result[key] = 0
	}

	result[key]++
}

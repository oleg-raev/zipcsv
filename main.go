package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"log"
	"strings"
	"csvaggr/zipcsv"
)

/*
дата/время, //2017-07-01T09:28:22
id вестибюля,
номер билета,
UID носителя,
Тип билета,
Тип прохода ( 0 - проходы, 1 - внешние пересадки, -1 - внутренние пересадки),
ѕор порядковый номер поездки по билету (если будет проставлен в системе),
количество оставшихся поездок по билету (если будет проставлено в системе)
 */




func main() {
	result := make(map[string]int64)

	dir := getCurrentDir()
	files := listFilesOfDir(dir)

	fmt.Printf("Hello world. Current dir is: %q\n", dir)
	fmt.Printf("ZIP files in the directory: %v\n", files)
	if len(files) < 1 {
		fmt.Println("Have no found ZIP files in the directory")
		os.Exit(0)
	}

	rows, errs := zipcsv.ProcessFiles(files)

	counter :=

	loop:
	for {
		select {
		case row, ok := <- rows:
			if !ok {
				break loop
			}
			processRow(row, &result)
		case err, ok := <- errs:
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

func processRow(row string, result *map[string]int64) {
	data := strings.Split(row, ";")
	if len(data) < 1 || len(data[0]) < 13 {
		return
	}

	key := data[0][11:13]

	if _, ok := (*result)[key]; !ok {
		(*result)[key] = 0
	}

	(*result)[key]++
}


func getCurrentDir() string {
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		log.Fatal(err)
		return ""
	} else {
		return dir
	}
	return ""
}

func listFilesOfDir(dir string) []string {
	var result []string

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() || !zipcsv.IsZIP(file.Name()) {
			continue
		}
		result = append(result, dir + string(os.PathSeparator) + file.Name())
	}

	return result
}
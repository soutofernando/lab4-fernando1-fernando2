package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

// sum all bytes of a file
func sum(data []byte) (int, error) {
    _sum := 0

	if err != nil {
		return 0, err
	}

	for _, b := range data {
		_sum += int(b)
	} 

	return _sum, nil
}

func worker (filePath string, results chan<- fileSumResult){
	data, err := readFile(filePath)
	if err != nil {
		results <- fileSumResult{filePath: filePath, sum: 0, err : err}
		return
	}
	_sum := sum(data)
	results <- fileSumResult{filePath: filePath, sum: _sum, err : nil}
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64
	sums := make(map[int][]string)
	for _, path := range os.Args[1:] {
		_sum, err := sum(path)

		if err != nil {
			continue
		}

		totalSum += int64(_sum)

		sums[_sum] = append(sums[_sum], path)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}
}

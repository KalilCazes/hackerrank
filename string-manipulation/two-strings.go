package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func twoStrings(s1 string, s2 string) string {

	var result string

	firstFrequency := make(map[rune]int)
	secondFrequency := make(map[rune]int)

	for _, v := range s1 {
		firstFrequency[v]++
	}

	for _, v := range s2 {
		secondFrequency[v]++
	}

	for k, _ := range firstFrequency {

		if _, ok := secondFrequency[k]; ok {
			result = "YES"
			break
		} else {
			result = "NO"
		}
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s1 := readLine(reader)

		s2 := readLine(reader)

		result := twoStrings(s1, s2)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

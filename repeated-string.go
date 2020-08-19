package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func repeatedString(s string, n int64) int64 {
	var groupCount int64
	var remainingCount int64
	var groups = int(n) / len(s)
	var remaining = int(n) % len(s)
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			groupCount++

			if i < remaining {
				remainingCount++
			}
		}
	}

	return groupCount*int64(groups) + remainingCount
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	fmt.Fprintf(writer, "%d\n", result)

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

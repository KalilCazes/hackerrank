package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func makeAnagram(a string, b string) int32 {
	var diff int32
	m := make(map[rune]int, 26)

	for _, v := range a {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	for _, v := range b {
		if _, ok := m[v]; ok {
			m[v]--

			if m[v] == 0 {
				delete(m, v)
			}
		} else {
			diff++
		}
	}

	for _, v := range m {
		diff += int32(v)
	}
	return diff

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

	fmt.Fprintf(writer, "%d\n", res)

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

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sliceToMap(arr []int32, m map[int]int) map[int]int {
	for i, v := range arr {
		m[int(v)] = i
	}

	return m
}

func minimumSwaps(arr []int32) int32 {
	var highest int
	var swap int32
	highest = len(arr)
	m := make(map[int]int)
	m = sliceToMap(arr, m)

	for i := 0; i < len(arr); i++ {

		if m[highest] != highest-1 {

			tmp := m[highest]

			m[highest] = highest - 1

			m[int(arr[highest-1])] = tmp

			arr[highest-1], arr[tmp] = arr[tmp], arr[highest-1]
			swap++
		}
		highest--
	}

	return swap
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

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

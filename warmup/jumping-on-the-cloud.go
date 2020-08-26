package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func jumpingOnClouds(c []int32) int32 {
	var jump int32
	var last, secondLast int

	last, secondLast = len(c)-1, len(c)-2

	for i := 0; i < len(c); i++ {

		if i < secondLast && c[i+2] == 0 {
			jump++
			i++

		} else if i < last && c[i+1] == 0 {
			jump++

		}
	}

	return jump
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

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := jumpingOnClouds(c)

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

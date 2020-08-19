	package main

	import (
		"bufio"
		"fmt"
		"io"
		"os"
		"strconv"
		"strings"
	)

	type num struct {
		value int32
		count int32
	}

	func queueToStruct(q []int32, n []num) []num {
		for _, v := range q {

			n = append(n, num{value: v, count: 0})
		}

		return n
	}
	func minimumBribes(q []int32) {
		var sum int32
		var num []num

		n := queueToStruct(q, num)

		for k := 0; k < len(n)-1; k++ {

			if n[k].value != int32(k)+1{
				for i := len(n) - 1; i > 0; i-- {

					if n[i].value < n[i-1].value {

						n[i], n[i-1] = n[i-1], n[i]	
						n[i].count++

						if n[i].count > 2 || n[i-1].count > 2{
							fmt.Println("Too chaotic")
							return
						}	
						sum++
					}
				}
			}
		}
		fmt.Println(sum)
	}

	func main() {
		reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

		tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		t := int32(tTemp)

		for tItr := 0; tItr < int(t); tItr++ {
			nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
			checkError(err)
			n := int32(nTemp)

			qTemp := strings.Split(readLine(reader), " ")

			var q []int32

			for i := 0; i < int(n); i++ {
				qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
				checkError(err)
				qItem := int32(qItemTemp)
				q = append(q, qItem)
			}

			minimumBribes(q)
		}
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

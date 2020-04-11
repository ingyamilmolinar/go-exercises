package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	countHash := make(map[int32]int32)
	freqHash := make(map[int32][]int32)
	var res []int32
	for _, query := range queries {
		command := query[0]
		element := query[1]
		switch command {
		case 1:
			freqHash[countHash[element]] = remove(freqHash[countHash[element]], element)
			countHash[element]++
			freqHash[countHash[element]] = append(freqHash[countHash[element]], element)

		case 2:
			if countHash[element] > 0 {
				freqHash[countHash[element]] = remove(freqHash[countHash[element]], element)
				countHash[element]--
				freqHash[countHash[element]] = append(freqHash[countHash[element]], element)
			}
		case 3:
			if len(freqHash[element]) > 0 {
				res = append(res, int32(1))
			} else {
				res = append(res, int32(0))
			}
		default:
			return res
		}
	}
	return res
}

// innefficient
func remove(s []int32, elem int32) []int32 {
	length := len(s)
	if length == 0 {
		return s
	}
	for i, n := range s {
		if n == elem {
			s[i] = s[length-1]
			return s[:length-1]
		}
	}
	return s
}

func main() {
	file, err := os.Open(os.Getenv("INPUT_PATH"))
	checkError(err)

	defer file.Close()

	reader := bufio.NewReader(file)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := freqQuery(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, err := reader.ReadString('\n')
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

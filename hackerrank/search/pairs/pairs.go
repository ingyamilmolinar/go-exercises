package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the pairs function below.
func pairs(k int32, arr []int32) int32 {
	hash := make(map[int32][]int, len(arr))
	for i, value := range arr {
		hash[value] = append(hash[value], i)
	}
	count := int32(0)
	for i, value := range arr {
		for _, idx := range hash[value+k] {
			if idx > i {
				count++
			}
		}
		for _, idx := range hash[value-k] {
			if idx > i {
				count++
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open(os.Getenv("INPUT_PATH"))
	checkError(err)

	defer file.Close()

	reader := bufio.NewReader(file)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := pairs(k, arr)

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

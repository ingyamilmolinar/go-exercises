package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func quickSort(arr []int32) []int32 {
	if len(arr) < 2 {
		return arr
	}
	pivotIdx := rand.Intn(len(arr))
	pivot := arr[pivotIdx]
	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[i] >= pivot && arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			for i < len(arr)-1 && arr[i] == arr[j] && arr[i] == pivot {
				i++
			}
		} else if arr[i] < pivot {
			i++
		} else if arr[j] > pivot {
			j--
		}
	}

	return append(quickSort(arr[:i]), quickSort(arr[i:])...)
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

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < len(arrTemp); i++ {
		tmpInt, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		tmpInt32 := int32(tmpInt)
		arr = append(arr, tmpInt32)
	}

	result := quickSort(arr)

	fmt.Fprintf(writer, "%d\n", result)

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

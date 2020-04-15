package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func mergeSort(arr []int32) []int32 {
	length := len(arr)
	if length == 1 {
		return arr
	}
	if length == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
			return arr
		}
		return arr
	}
	arr1, arr2 := arr[:int(length/2)], arr[int(length/2):]
	return merge(mergeSort(arr1), mergeSort(arr2))
}

func merge(arr1 []int32, arr2 []int32) []int32 {
	result := make([]int32, len(arr1)+len(arr2))
	var i, x, y int
	for i < len(result) {
		if x < len(arr1) && y < len(arr2) {
			if arr1[x] < arr2[y] {
				result[i] = arr1[x]
				x++
			} else if arr1[x] > arr2[y] {
				result[i] = arr2[y]
				y++
			} else {
				result[i], result[i+1] = arr1[x], arr2[y]
				x++
				y++
				i++
			}
		} else if x < len(arr1) {
			result[i] = arr1[x]
			x++
		} else if y < len(arr2) {
			result[i] = arr2[y]
			y++
		}
		i++
	}
	return result
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

	result := mergeSort(arr)

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

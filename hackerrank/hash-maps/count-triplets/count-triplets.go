package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	// base case
	length := len(arr)
	if length < 3 {
		return 0
	}

	// sort
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	hash := make(map[int64]int64)
	count := int64(0)
	min, max := arr[0], arr[length-1]

	// Special case
	if r == 1 {
		// cluster in the hash
		for _, n := range arr {
			hash[n]++
		}
		count := int64(0)
		for _, c := range hash {
			count += getCombinations(c)
		}
		return count
	}

	// Normal case
	i := length - 1
	for i >= 0 {
		if arr[i]/r >= min && arr[i]%r == 0 {
			if arr[i]/r/r >= min && arr[i]%r%r == 0 {
				hash[arr[i]]++
			} else if _, ok := hash[arr[i]*r]; ok && arr[i]*r <= max {
				hash[arr[i]]++
			}
		}
		if _, ok := hash[arr[i]*r]; ok && arr[i]*r <= max {
			multiplier := hash[arr[i]*r]
			if _, ok := hash[arr[i]*r*r]; ok && arr[i]*r*r <= max {
				count += (hash[arr[i]*r*r] * multiplier)
			}
		}
		i--
	}

	return count
}

// This can be memoized
func getCombinations(n int64) int64 {
	if n < 3 {
		return 0
	} else if n == 3 {
		return 1
	}
	count := int64(0)
	for i := int64(2); i < n; i++ {
		count += n - i
	}
	return count + getCombinations(n-1)
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

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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

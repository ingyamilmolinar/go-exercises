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

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	hashMap := make(map[string]int32)
	length := len(s)
	for i, fChar := range s {
		hashMap[string(fChar)]++
		for j := range s {
			if j <= i || j-i == length-1 {
				continue
			}
			sorted := sortString(s[i : j+1])
			hashMap[sorted]++
		}
	}
	count := int32(0)
	for k, v := range hashMap {
		if v == 1 {
			continue
		}
		if len(k) == 1 {
			count += computeUntil(v)
			continue
		}
		count += computeUntil(v)
	}
	return count
}

func computeUntil(until int32) int32 {
	i := int32(0)
	result := int32(0)
	for i < until {
		result = i + result
		i++
	}
	return result
}

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
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

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

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

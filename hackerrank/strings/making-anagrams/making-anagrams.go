package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the makeAnagram function below.
func makeAnagram(a string, b string) int32 {
	hashA := make(map[rune]int32)
	for _, charA := range a {
		hashA[charA]++
	}
	hashB := make(map[rune]int32)
	for _, charB := range b {
		hashB[charB]++
	}
	count := int32(0)
	for charA, countA := range hashA {
		if _, found := hashB[charA]; found {
			if countA == hashB[charA] {
				continue
			}
			if countA-hashB[charA] > 0 {
				count += int32(countA - hashB[charA])
			}
			continue
		}
		count += countA
	}
	for charB, countB := range hashB {
		if _, found := hashA[charB]; found {
			if countB == hashA[charB] {
				continue
			}
			if countB-hashA[charB] > 0 {
				count += int32(countB - hashA[charB])
			}
			continue
		}
		count += countB
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

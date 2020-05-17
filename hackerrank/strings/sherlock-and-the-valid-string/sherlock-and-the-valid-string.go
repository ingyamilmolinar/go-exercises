package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

// Complete the isValid function below.
func isValid(s string) string {
	hash := make(map[rune]int)
	for _, char := range s {
		hash[char]++
	}
	freqs := make(map[int]int)
	maxFreq := 0
	minFreq := int32(2147483647)
	for _, count := range hash {
		freqs[count]++
		if count > maxFreq {
			maxFreq = count
		}
		if int32(count) < minFreq {
			minFreq = int32(count)
		}
	}
	if len(freqs) == 0 || len(freqs) == 1 {
		return "YES"
	} else if len(freqs) == 2 && ((freqs[maxFreq] == 1 && freqs[maxFreq-1] != 0) || (freqs[int(minFreq)] == 1 && (freqs[int(minFreq)+1] != 0 || minFreq == 1))) {
		return "YES"
	}
	return "NO"
}

func main() {
		file, err := os.Open(os.Getenv("INPUT_PATH"))
		checkError(err)

		defer file.Close()

		reader := bufio.NewReader(file)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    result := isValid(s)

    fmt.Fprintf(writer, "%s\n", result)

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

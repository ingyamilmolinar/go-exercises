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

type triplet struct {
	a int32
	b int32
	c int32
}

// Complete the triplets function below.
func triplets(a []int32, b []int32, c []int32) int64 {
	sortInt32(a)
	sortInt32(b)
	sortInt32(c)
	a = makeUnique(a)
	b = makeUnique(b)
	c = makeUnique(c)
	count := int64(0)
	var startA, startB, startC, endA, endC int
	for startB < len(b) {
		if a[endA] <= b[startB] && b[startB] >= c[endC] {
			for endA < len(a) && a[endA] <= b[startB] {
				endA++
			}
			endA--
			for endC < len(c) && c[endC] <= b[startB] {
				endC++
			}
			endC--
			count += int64(((endA - startA) + 1) * ((endC - startC) + 1))
			startB++
		} else if b[startB] < a[startA] || b[startB] < c[startC] {
			startB++
		}
	}
	return count
}

func sortInt32(a []int32) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] <= a[j]
	})
}

func makeUnique(a []int32) []int32 {
	var r []int32
	prevElem := a[0]
	r = append(r, prevElem)
	for _, elem := range a[1:] {
		if prevElem == elem {
			continue
		}
		r = append(r, elem)
		prevElem = elem
	}
	return r
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

	lenaLenbLenc := strings.Split(readLine(reader), " ")

	lenaTemp, err := strconv.ParseInt(lenaLenbLenc[0], 10, 64)
	checkError(err)
	lena := int32(lenaTemp)

	lenbTemp, err := strconv.ParseInt(lenaLenbLenc[1], 10, 64)
	checkError(err)
	lenb := int32(lenbTemp)

	lencTemp, err := strconv.ParseInt(lenaLenbLenc[2], 10, 64)
	checkError(err)
	lenc := int32(lencTemp)

	arraTemp := strings.Split(readLine(reader), " ")

	var arra []int32

	for i := 0; i < int(lena); i++ {
		arraItemTemp, err := strconv.ParseInt(arraTemp[i], 10, 64)
		checkError(err)
		arraItem := int32(arraItemTemp)
		arra = append(arra, arraItem)
	}

	arrbTemp := strings.Split(readLine(reader), " ")

	var arrb []int32

	for i := 0; i < int(lenb); i++ {
		arrbItemTemp, err := strconv.ParseInt(arrbTemp[i], 10, 64)
		checkError(err)
		arrbItem := int32(arrbItemTemp)
		arrb = append(arrb, arrbItem)
	}

	arrcTemp := strings.Split(readLine(reader), " ")

	var arrc []int32

	for i := 0; i < int(lenc); i++ {
		arrcItemTemp, err := strconv.ParseInt(arrcTemp[i], 10, 64)
		checkError(err)
		arrcItem := int32(arrcItemTemp)
		arrc = append(arrc, arrcItem)
	}

	ans := triplets(arra, arrb, arrc)

	fmt.Fprintf(writer, "%d\n", ans)

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

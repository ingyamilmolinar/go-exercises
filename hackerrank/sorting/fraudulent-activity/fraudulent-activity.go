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

// Complete the activityNotifications function below.
func activityNotifications(expenditure []int32, d int32) int32 {
	length := int32(len(expenditure))
	if length <= d {
		return 0
	}
	notifications := int32(0)
	trail := make([]int32, d)
	copy(trail, expenditure[0:d])
	sortSlice(trail)
	for i := d; i < length; i++ {
		median := getMedian(trail, d)
		if expenditure[i] >= 2*median {
			notifications++
		}
		if i+1 < length {
			trail = insertSorted(trail, d, expenditure[i-d], expenditure[i])
		}
	}
	return notifications
}

func sortSlice(a []int32) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
}

func insertSorted(a []int32, l int32, old int32, new int32) []int32 {
	op := sort.Search(int(l), func(i int) bool { return a[i] >= old })
	np := sort.Search(int(l), func(i int) bool { return a[i] >= new })
	a[op] = new

	// bubble up/down depending on value
	if op < np {
		for i := op; i < np-1; i++ {
			a[i], a[i+1] = a[i+1], a[i]
		}
	} else if op > np {
		for i := op; i > np; i-- {
			a[i], a[i-1] = a[i-1], a[i]
		}
	}
	return a
}

func getMedian(a []int32, l int32) int32 {
	if l%2 == 0 {
		return (a[l/2] + a[(l/2)+1]) / 2
	}
	return a[int(l/2)]
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

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(readLine(reader), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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

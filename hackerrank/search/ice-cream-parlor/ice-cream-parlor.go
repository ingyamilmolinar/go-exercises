package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the whatFlavors function below.
func whatFlavors(cost []int32, money int32) {
	hash := make(map[int32][]int, len(cost))
	for i, c := range cost {
		hash[c] = append(hash[c], i)
	}
	for i, c := range cost {
		for _, idx := range hash[money-c] {
			if i < idx {
				fmt.Println(strconv.Itoa(i+1) + " " + strconv.Itoa(idx+1))
				return
			}
		}
	}
}

func main() {
	file, err := os.Open(os.Getenv("INPUT_PATH"))
	checkError(err)

	defer file.Close()

	reader := bufio.NewReader(file)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		moneyTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		money := int32(moneyTemp)

		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		costTemp := strings.Split(readLine(reader), " ")

		var cost []int32

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
	}
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

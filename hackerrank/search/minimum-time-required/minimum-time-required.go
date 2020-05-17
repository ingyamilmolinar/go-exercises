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

// Complete the minTime function below.
func minTime(machines []int64, goal int64) int64 {
	sort.Slice(machines, func(i, j int) bool {
		return machines[i] <= machines[j]
	})
	floatMaxDays := float64(machines[len(machines)-1]*goal) / float64(len(machines))
	maxDays := int64(floatMaxDays)
	if maxDays%machines[len(machines)-1] != 0 {
		maxDays += machines[len(machines)-1]
	}
	floatMinDays := float64(machines[0]*goal) / float64(len(machines))
	minDays := int64(floatMinDays)
	middleDays := float64(maxDays+minDays) / 2.0
	var estimatedGoal int64
	for {
		estimatedGoal = computeProductionFromDays(machines[:], int64(middleDays))
		if estimatedGoal > goal {
			maxDays = int64(middleDays)
		} else if estimatedGoal < goal {
			minDays = int64(middleDays)
		} else {
			// search for min ocurrence
			prevEstimatedGoal := estimatedGoal
			newMiddleDays := int64(middleDays)
			for prevEstimatedGoal == estimatedGoal {
				prevEstimatedGoal = estimatedGoal
				newMiddleDays = newMiddleDays - 1
				estimatedGoal = computeProductionFromDays(machines[:], newMiddleDays)
			}
			return newMiddleDays + 1
		}
		middleDays = float64(maxDays+minDays) / 2.0
	}
}

func computeProductionFromDays(machines []int64, days int64) int64 {
	total := int64(0)
	for _, machine := range machines {
		total += (days / machine)
	}
	return total
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

	nGoal := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nGoal[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	goal, err := strconv.ParseInt(nGoal[1], 10, 64)
	checkError(err)

	machinesTemp := strings.Split(readLine(reader), " ")

	var machines []int64

	for i := 0; i < int(n); i++ {
		machinesItem, err := strconv.ParseInt(machinesTemp[i], 10, 64)
		checkError(err)
		machines = append(machines, machinesItem)
	}

	ans := minTime(machines, goal)

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

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type pair struct {
	fst string
	snd string
}

func (this *pair) first() string {
	return this.fst
}

func (this *pair) second() string {
	return this.snd
}

type stackMemory struct {
	p pair
	c int32
}

type stack struct {
	len   int
	stack []stackMemory
	//set   map[pair]bool
}

func newStack() stack {
	var stack stack
	//stack.set = make(map[pair]bool)
	return stack
}

func (this *stack) push(sm stackMemory) {
	this.stack = append(this.stack, sm)
	this.len++
	//this.set[sm.p] = true
}

func (this *stack) pop() (stackMemory, error) {
	if this.len > 0 {
		popped := this.stack[this.len-1]
		this.stack = this.stack[:this.len-1]
		this.len--
		//this.set[popped.p] = false
		return popped, nil
	}
	return stackMemory{}, fmt.Errorf("Cannot pop empty stack")
}

func (this *stack) isEmpty() bool {
	return this.len <= 0
}

// Complete the commonChild function below.
func commonChild(s1 string, s2 string) int32 {
	memo := make(map[pair]int32)
	return countSubstringsIter(s1, s2, len(s1), len(s2), memo)
	//return countSubstringsRec(s1, s2, len(s1), len(s2), memo)
}

// TODO: Memoization is missing
func countSubstringsIter(s1 string, s2 string, len1 int, len2 int, memo map[pair]int32) int32 {
	var count int32
	stack := newStack()
	stack.push(stackMemory{pair{s1, s2}, int32(0)})
	for !stack.isEmpty() {
		popped, _ := stack.pop()
		if popped.p.first() == "" || popped.p.second() == "" {
			continue
		}
		if popped.p.first()[0] == popped.p.second()[0] {
			stack.push(stackMemory{pair{popped.p.first()[1:], popped.p.second()[1:]}, popped.c + 1})
			if popped.c+1 > count {
				count = popped.c + 1
			}
		} else {
			stack.push(stackMemory{pair{popped.p.first()[1:], popped.p.second()}, popped.c})
			stack.push(stackMemory{pair{popped.p.first(), popped.p.second()[1:]}, popped.c})
		}
	}
	return count
}

func countSubstringsRec(s1 string, s2 string, len1 int, len2 int, memo map[pair]int32) int32 {
	memoPair := pair{s1, s2}
	if _, ok := memo[memoPair]; ok {
		return memo[memoPair]
	}
	if len1 == 0 || len2 == 0 {
		memo[memoPair] = 0
		return memo[memoPair]
	}
	if s1[0] == s2[0] {
		memo[memoPair] = 1 + countSubstringsRec(s1[1:], s2[1:], len1-1, len2-1, memo)
		return memo[memoPair]
	}
	memo[memoPair] = max(countSubstringsRec(s1[1:], s2[0:], len1-1, len2, memo), countSubstringsRec(s1[0:], s2[1:], len1, len2-1, memo))
	return memo[memoPair]
}

func max(n1 int32, n2 int32) int32 {
	if n1 > n2 {
		return n1
	}
	return n2
}

func countSubstrings1(s string, length int, prevIdx int, hash map[rune][]int) int32 {
	if length == 0 {
		return 0
	}
	maxCount := int32(0)
	for i, char := range s {
		if idxs, ok := hash[char]; ok {
			for _, idx := range idxs {
				if idx > prevIdx || prevIdx == -1 {
					tmpCount := 1 + countSubstrings1(s[i+1:], length-1, idx, hash)
					if tmpCount > maxCount {
						maxCount = tmpCount
					}
				}
			}
		}
	}
	return maxCount
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

	s1 := readLine(reader)

	s2 := readLine(reader)

	result := commonChild(s1, s2)

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

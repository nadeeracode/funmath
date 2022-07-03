package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputString string
	fmt.Printf("Input a number ")
	reader := bufio.NewReader(os.Stdin)
	inputString, _ = reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.Trim(inputString, "\n"))
	fmt.Printf("First %d numbers of Fibonacci Sequence ", n)
	var s1 int = 1
	var s2 int = 1
	var s3 int = 0
	fmt.Printf("%d ", s1)
	fmt.Printf("%d ", s2)
	for i := 2; i < n; i++ {
		s3 = s2 + s1
		fmt.Printf("%d ", s3)
		s1 = s2
		s2 = s3
	}
}

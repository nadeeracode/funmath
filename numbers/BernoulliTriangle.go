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
	fmt.Printf("Input another number ")
	inputString, _ = reader.ReadString('\n')
	k, _ := strconv.Atoi(strings.Trim(inputString, "\n"))
	fmt.Printf("First %d %d numbers of Bernoulli Triangle \n ", n, k)
	var B [][]int = make([][]int, n)
	for i := range B {
		B[i] = make([]int, k)
	}
	B[0][0] = 1
	for i := 1; i < n; i++ {
		for j := 0; j < k; j++ {
			if j == 0 {
				B[i][j] = 1
			} else if i == j {
				B[i][j] = 2 * B[i-1][j-1]
			} else {
				B[i][j] = B[i-1][j] + B[i-1][j-1]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			if B[i][j] == 0 {
				continue
			}
			fmt.Printf("%d ", B[i][j])
		}
		fmt.Printf("\n ")
	}
	//fmt.Println(B)
}

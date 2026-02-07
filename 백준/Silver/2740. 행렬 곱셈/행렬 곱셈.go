package main

import (
	"bufio"
	"os"
    "fmt"
)

var reader *bufio.Reader
var writer *bufio.Writer

func main() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()
    var n, m, k int
	fmt.Fscanf(reader, "%d %d \n", &n, &m)
	var A = make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscanf(reader, "%d ", &A[i][j])
		}
	}
	fmt.Fscanf(reader, "%d %d \n", &m, &k)
	var B = make([][]int, m)
	for i := 0; i < m; i++ {
		B[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscanf(reader, "%d ", &B[i][j])
		}
	}
	var res = make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, k)
		for j := 0; j < k; j++ {
			for k := 0; k < m; k++ {
				res[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			fmt.Fprintf(writer, "%d ", res[i][j])
		}
		fmt.Fprintf(writer, "\n")
	}
}

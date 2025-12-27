package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcb(a, b int) int {
	for a > 0 {
		a, b = b%a, a
	}
	return b
}
func lcm(a, b int) int {
	return a * b / gcb(a, b)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, a, b int
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		fmt.Fprintf(writer, "%d\n", lcm(a, b))
	}
}

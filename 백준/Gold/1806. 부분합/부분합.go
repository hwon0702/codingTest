package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	var s int64
	fmt.Fscan(reader, &n, &s)

	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	var sum int64 = 0
	left := 0
	res := n + 1

	for right := 0; right < n; right++ {
		sum += arr[right]

		for sum >= s {
			if right-left+1 < res {
				res = right - left + 1
			}
			sum -= arr[left]
			left++
		}
	}

	if res == n+1 {
		fmt.Fprintln(writer, 0)
	} else {
		fmt.Fprintln(writer, res)
	}
}

package completebinarytree

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	total, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input number")
	}

	nodes := make([]int, total+1)

	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	for i := 0; i < total; i++ {
		node, err := strconv.Atoi(inputs[i])
		if err != nil {
			fmt.Println("Input number")
		}
		nodes[i+1] = node
	}

	for i, node := range nodes {
		if i == 0 {
			continue
		}
		fmt.Printf("node: %d key: %d ", i, node)
		p := parent(i)
		if p > 0 {
			fmt.Printf("parent: %d ", nodes[p])
		}
		left := left(i)
		if left <= total {
			fmt.Printf("left: %d ", nodes[left])
		}
		right := right(i)
		if right <= total {
			fmt.Printf("right: %d ", nodes[right])
		}
		fmt.Print("\n")

	}

}

func parent(n int) int {
	return n / 2
}

func left(n int) int {
	return 2 * n
}

func right(n int) int {
	return 2*n + 1
}

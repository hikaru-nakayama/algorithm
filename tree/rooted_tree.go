package tree

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Parent int64 // ノードの親
	Left   int64 // ノードの最も左の子
	Right  int64 // ノードのすぐ右の兄弟
}

type Tree struct {
	Nodes []Node
}

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	for _, input := range inputs {
		fmt.Println(input)
	}
}

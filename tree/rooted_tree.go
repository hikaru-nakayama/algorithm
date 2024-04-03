package tree

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Parent string // ノードの親
	Left   string // ノードの最も左の子
	Right  string // ノードのすぐ右の兄弟
}

type Tree struct {
	Nodes []Node
}

func (n *Node) String(id int) {
	fmt.Printf("Self: %d Parent: %s Left %s Right %s\n", id, n.Parent, n.Left, n.Right)
}

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	totalNodeNum, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("数字を入力せよ")
	}
	tree := Tree{
		Nodes: make([]Node, totalNodeNum),
	}
	for i := 0; i < totalNodeNum; i++ {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
		childs := inputs[1:]
		tree.Nodes[i].Left = childs[0]
		for _, c := range childs {
			ch, err := strconv.Atoi(c)
			if err != nil {
				fmt.Println("数字を入力せよ")
			}
			tree.Nodes[ch].Parent = strconv.Itoa(i)
			if ch+1 < len(childs) {
				tree.Nodes[ch].Right = childs[ch+1]
			}
		}
	}
	for i, n := range tree.Nodes {
		n.String(i)
	}
}

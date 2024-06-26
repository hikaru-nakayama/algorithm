package rooted_tree

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
	fmt.Printf("Self: %d Parent: %s Left %s Right %s ", id, n.Parent, n.Left, n.Right)
}

func Start() {
	var Nil = "-1"
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
	for i := range tree.Nodes {
		tree.Nodes[i].Parent = Nil
		tree.Nodes[i].Right = Nil
	}
	for i := 0; i < totalNodeNum; i++ {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
		if inputs[1] == "0" {
			tree.Nodes[i].Left = Nil
			continue
		}
		childs := inputs[2:]
		tree.Nodes[i].Left = childs[0]
		for j, c := range childs {
			ch, err := strconv.Atoi(c)
			if err != nil {
				fmt.Println("数字を入力せよ")
			}
			tree.Nodes[ch].Parent = strconv.Itoa(i)
			if j+1 < len(childs) {
				tree.Nodes[ch].Right = childs[j+1]
			}
		}
	}
	deepList := tree.getDeep(0, 0, make([]int, totalNodeNum))
	for i, n := range tree.Nodes {
		n.String(i)
		fmt.Printf("deep: %d ", deepList[i])
		fmt.Printf("type: %s\n", n.getType())
	}
}

func (t *Tree) getDeep(u int, p int, d []int) []int {
	d[u] = p
	if t.Nodes[u].Right != "-1" {
		right, err := strconv.Atoi(t.Nodes[u].Right)
		if err != nil {
			fmt.Println("数字を入力せよ")
		}
		t.getDeep(right, p, d)
	}
	if t.Nodes[u].Left != "-1" {
		right, err := strconv.Atoi(t.Nodes[u].Left)
		if err != nil {
			fmt.Println("数字を入力せよ")
		}
		t.getDeep(right, p+1, d)
	}
	return d
}

func (n *Node) getType() string {
	if n.Left == "-1" {
		return "leaf"
	} else if n.Parent == "-1" {
		return "root"
	} else {
		return "internal"
	}
}

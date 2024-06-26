package treework

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Parent string // 親
	Left   string // 左の子
	Right  string // 右の子
}

type Tree struct {
	Nodes []Node
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
		tree.Nodes[i].Left = "-1"
		tree.Nodes[i].Right = "-1"
		tree.Nodes[i].Parent = "-1"
	}

	for i := 0; i < totalNodeNum; i++ {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
		id, err := strconv.Atoi(inputs[0])
		if err != nil {
			fmt.Println("数字を入力せよ")
		}
		tree.Nodes[id].Left = inputs[1]
		tree.Nodes[id].Right = inputs[2]
		left, err := strconv.Atoi(inputs[1])
		if err != nil {
			fmt.Println("数字を入力せよ")
		}
		right, err := strconv.Atoi(inputs[2])
		if err != nil {
			fmt.Println("数字を入力せよ")
		}
		if left == -1 || right == -1 {
			continue
		}
		tree.Nodes[left].Parent = inputs[0]
		tree.Nodes[right].Parent = inputs[0]
	}

	for i := 0; i < totalNodeNum; i++ {
		node := tree.Nodes[i]
		fmt.Printf("Node: %d Parent: %s Left: %s Right: %s\n", i, node.Parent, node.Left, node.Right)
	}

	fmt.Print("\n")
	fmt.Println("Preorder")
	tree.preOrder(0)
	fmt.Print("\n")
	fmt.Println("Inorder")
	tree.inOrder(0)
	fmt.Print("\n")
	fmt.Println("Postorder")
	tree.postOrder(0)
	fmt.Print("\n")

}

func (t *Tree) preOrder(u int) {
	fmt.Printf("%d ", u)
	if t.Nodes[u].Left != "-1" {
		left, err := strconv.Atoi(t.Nodes[u].Left)
		if err != nil {
			fmt.Println("数字を入力してください")
		}
		t.preOrder(left)
	}
	if t.Nodes[u].Right != "-1" {
		right, err := strconv.Atoi(t.Nodes[u].Right)
		if err != nil {
			fmt.Println("数字を入力してください")
		}
		t.preOrder(right)
	}
}

func (t *Tree) inOrder(u int) {
	if t.Nodes[u].Left != "-1" {
		left, err := strconv.Atoi(t.Nodes[u].Left)
		if err != nil {
			fmt.Println("数字を入力してください")
		}
		t.inOrder(left)
	}
	fmt.Printf("%d ", u)
	if t.Nodes[u].Right != "-1" {
		right, err := strconv.Atoi(t.Nodes[u].Right)
		if err != nil {
			fmt.Println("数字を入力してください")
		}
		t.inOrder(right)
	}
}

func (t *Tree) postOrder(u int) {
	if t.Nodes[u].Left != "-1" {
		left, err := strconv.Atoi(t.Nodes[u].Left)
		if err != nil {
			fmt.Println("数字を入力してください")
		}
		t.postOrder(left)
	}
	if t.Nodes[u].Right != "-1" {
		right, err := strconv.Atoi(t.Nodes[u].Right)
		if err != nil {
			fmt.Println("数字を入力してください")
		}
		t.postOrder(right)
	}
	fmt.Printf("%d ", u)
}

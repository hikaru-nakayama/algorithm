package binarysearch

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
}

type Tree struct {
	Nodes []Node
}

var Nil = -1

func Insert() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	totalNodeNum, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input Number")
	}

	tree := Tree{Nodes: make([]Node, totalNodeNum)}
	for i := 0; i < totalNodeNum; i++ {
		tree.Nodes[i].Key = Nil
		tree.Nodes[i].Right = &Node{Key: Nil}
		tree.Nodes[i].Left = &Node{Key: Nil}
	}

	for i := 0; i < totalNodeNum; i++ {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
		key, err := strconv.Atoi(inputs[1])
		if err != nil {
			fmt.Println("Input Number")
		}

		if i == 0 {
			tree.Nodes[i].Key = key
			tree.Nodes[i].Parent = &Node{Key: Nil}
			continue
		}

		cur := &tree.Nodes[0] // index 0 を rootNode とする.
		parent := &tree.Nodes[0]

		for cur.Key != Nil {
			parent = cur
			if cur.Key < key {
				cur = cur.Right
			} else {
				cur = cur.Left
			}
		}

		cur.Parent = parent
		if cur.Parent.Key < key {
			cur.Parent.Right = cur
		} else if cur.Parent.Key > key {
			cur.Parent.Left = cur
		}
		cur.Key = key
		tree.Nodes[i] = *cur

	}

	for _, node := range tree.Nodes {
		fmt.Printf("Key: %d Parent: %d ", node.Key, node.Parent.Key)
	}
	fmt.Print("\n")
}

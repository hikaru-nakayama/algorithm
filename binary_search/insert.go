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
	}

	for i := 0; i < totalNodeNum; i++ {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")
		key, err := strconv.Atoi(inputs[1])
		if err != nil {
			fmt.Println("Input Number")
		}

		tree.Nodes[i] = Node{Key: key}

		cur := &tree.Nodes[0] // index 0 を rootNode とする.
		parent := cur

		for cur.Key != Nil {
			if cur.Key < key {
				parent = cur
				cur = cur.Right
			} else if cur.Key > key {
				parent = cur
				cur = cur.Left
			} else {
				break
			}
		}

		tree.Nodes[i].Parent = parent
		if cur.Parent.Key < key {
			tree.Nodes[i].Parent.Right = cur
		} else if cur.Parent.Key > key {
			tree.Nodes[i].Parent.Left = cur
		}

	}

	for _, node := range tree.Nodes {
		fmt.Printf("Parent: %d", node.Parent.Key)
	}
}

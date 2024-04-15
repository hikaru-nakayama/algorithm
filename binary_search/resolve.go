package binarysearch

import (
	"bufio"
	"errors"
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

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	totalNodeNum, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input Number")
	}

	tree := Tree{Nodes: make([]Node, totalNodeNum)}

	for i := 0; i < totalNodeNum; i++ {
		scanner.Scan()
		inputs := strings.Split(scanner.Text(), " ")

		key, err := strconv.Atoi(inputs[1])
		if err != nil {
			fmt.Println("Input Number")
		}

		if inputs[0] == "find" {
			tree.search(key)
			continue
		} else if inputs[0] == "delete" {
			tree.delete(key)
			continue
		}

		if i == 0 {
			tree.Nodes[i].Key = key
			tree.Nodes[i].Parent = &Node{Key: Nil}
			tree.Nodes[i].Right = &Node{Key: Nil}
			tree.Nodes[i].Left = &Node{Key: Nil}
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
		cur.Right = &Node{Key: Nil}
		cur.Left = &Node{Key: Nil}
		cur.Key = key
		if cur.Parent.Key < key {
			cur.Parent.Right = cur
		} else if cur.Parent.Key > key {
			cur.Parent.Left = cur
		}
		tree.Nodes[i] = *cur

	}
	tree.preOrder(tree.Nodes[0])
}

func (t *Tree) preOrder(n Node) {
	if n.Left.Key != Nil {
		t.preOrder(*n.Left)
	}
	fmt.Printf("%d ", n.Key)
	if n.Right.Key != Nil {
		t.preOrder(*n.Right)
	}
}

func (t *Tree) search(k int) {
	cur := t.Nodes[0]
	for cur.Key != Nil {
		if cur.Key < k {
			cur = *cur.Right
		} else if cur.Key > k {
			cur = *cur.Left
		} else {
			fmt.Println("yes")
			return
		}
	}
	fmt.Println("no")
}

func (t *Tree) searchNextNode(k int) (*Node, error) {
	cur := t.Nodes[0]
	for cur.Key != Nil {
		if cur.Key < k {
			cur = *cur.Right
		} else if cur.Key > k {
			cur = *cur.Left
		} else {
			return &cur, nil
		}
	}
	return &Node{}, errors.New("not found")
}

func (t *Tree) searchNode(k int) (*Node, error) {
	cur := t.Nodes[0]
	for cur.Key != Nil {
		if cur.Key < k {
			cur = *cur.Right
		} else if cur.Key > k {
			cur = *cur.Left
		} else {
			return &cur, nil
		}
	}
	return &Node{}, errors.New("not found")
}

func (t *Tree) delete(k int) {
	node, err := t.searchNode(k)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	if node.Left.Key == Nil {
		child := node.Right
		child.Parent = node.Parent
		if node.Parent.Left == node {
			node.Parent.Left = child
		} else if node.Parent.Right == node {
			node.Parent.Right = child
		}
	} else if node.Right.Key == Nil {
		child := node.Left
		child.Parent = node.Parent
		if node.Parent.Left == node {
			node.Parent.Left = child
		} else if node.Parent.Right == node {
			node.Parent.Right = child
		}
	} else {
		targetNode, err := t.searchNextNode(k)
		if err != nil {
			fmt.Printf("%s", err.Error())
		}
		if targetNode.Right.Key != Nil {
			targetNode.Parent.Left = targetNode.Right
			targetNode.Right.Parent = targetNode.Parent
		}
		node.Key = targetNode.Key
	}
}

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
	tree.inOrder(tree.Nodes[0])
	node, err := tree.searchNextNode(&tree.Nodes[0])
	if err != nil {
		fmt.Print("ss")
	}
	fmt.Print("\n")
	fmt.Printf("next node is %d\n", node.Key)
}

func (t *Tree) inOrder(n Node) {
	if n.Left.Key != Nil {
		t.inOrder(*n.Left)
	}
	fmt.Printf("%d ", n.Key)
	if n.Right.Key != Nil {
		t.inOrder(*n.Right)
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

func (t *Tree) searchNextNode(n *Node) (*Node, error) {
	if n.Right.Key != Nil {
		return getMinimum(n.Right), nil
	}

	parent := n.Parent
	for parent.Key != Nil && n != parent.Left {
		n = parent
		parent = n.Parent
	}

	return parent, nil
}

func getMinimum(n *Node) *Node {
	for n.Left.Key != Nil {
		n = n.Left
	}
	return n
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
		p := node.Parent
		child.Parent = p
		if p.Left.Key == node.Key {
			p.Left = child
		} else if p.Right.Key == node.Key {
			p.Right = child
		}
	} else if node.Right.Key == Nil {
		child := node.Left
		p := node.Parent
		child.Parent = p
		if p.Left.Key == node.Key {
			p.Left = child
		} else if p.Right.Key == node.Key {
			p.Right = child
		}
	} else {
		targetNode, err := t.searchNextNode(node)
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

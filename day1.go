package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// given two lists, i need to sort them and find the distance
	// between the ordered numbers (asc)

	// since i need to read the lists, i could take each number and insert it
	// into an already ordered list, finding where it needs to be placed with
	// binary search (?)

	// i could also do a dumb insertion, going through the list linearly until
	// n = 100 or something. for this a linked list would be useful

	input, _ := os.ReadFile("./sample-inputs/day1.txt")

	part1(string(input))
}

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) printAhead() {
	for currentNode := n; currentNode != nil; currentNode = currentNode.Next {
		fmt.Printf("%d > ", currentNode.Value)
	}
	fmt.Print("\n")
}

func part1(input string) {
	firstRoot := &Node{}
	secondRoot := &Node{}

	for _, line := range strings.Split(input, "\n") {
		lineParts := strings.Split(line, "   ")

		firstNumber, err := strconv.Atoi(lineParts[0])

		if err != nil {
			continue
		}

		secondNumber, err := strconv.Atoi(lineParts[1])

		if err != nil {
			continue
		}

		firstRoot = insertInList(firstNumber, firstRoot)
		secondRoot = insertInList(secondNumber, secondRoot)
	}

	// distanceSum := 0

	for ; firstRoot != nil; firstRoot = firstRoot.Next {

	}
}

// returns new root in case it changes
func insertInList(value int, root *Node) *Node {
	// four cases to consider
	// 1. if the root is empty, new value must go in root
	// 2. if the root is greater than the next number, new value must become
	// 		new root
	// 3. if the value is greater than the last number, it should be added after
	// 4. if the value is greater than current node and <= next node, it should
	// 		be included between them

	// case 1 (we asume that there are no 0's in the list)
	if root.Value == 0 {
		root.Value = value
		return root
	}

	// case 2
	if root.Value >= value {
		newRoot := Node{Value: value, Next: root}
		return &newRoot
	}

	// case 4
	currentNode := root

	for ; currentNode.Next != nil; currentNode = currentNode.Next {
		if currentNode.Value <= value && value <= currentNode.Next.Value {
			newNode := Node{Value: value, Next: currentNode.Next}
			currentNode.Next = &newNode
			return root
		}
	}

	// case 3
	currentNode.Next = &Node{Value: value}

	return root
}

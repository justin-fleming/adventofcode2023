package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name  string
	left  string
	right string
}

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	var firstNode Node

	var steps []string
	var nodes []Node
	nodeMap := make(map[string]Node)
	count := 1
	for scanner.Scan() {
		// TODO: iterate over lines of input here...
		if count == 1 {
			steps = strings.Split(scanner.Text(), "")
		}
		if count > 2 {
			lines := strings.Split(scanner.Text(), "=")
			nodeName := strings.TrimSpace(lines[0])
			lr := strings.Split(lines[1], ", ")
			n := Node{
				name:  nodeName,
				left:  strings.Trim(strings.TrimSpace(lr[0]), "()"),
				right: strings.Trim(strings.TrimSpace(lr[1]), "()"),
			}
			nodes = append(nodes, n)
			nodeMap[nodeName] = n
			if nodeName == "AAA" {
				firstNode = n
			}
		}
		count++
	}

	fmt.Println("Steps:", steps)
	fmt.Println("Nodes:", nodeMap)
	fmt.Println("FIRST NODE IS:", firstNode)

	endLoc := "ZZZ"
	stepCount := 1
	currentNode := firstNode
	fmt.Println("starting current node:", currentNode)
outerLoop:
	for {
		fmt.Println("outer loop..")
		for _, s := range steps {
			currentNode = followNode(currentNode, s, nodeMap)
			if currentNode.name == endLoc {
				break outerLoop
			}
			stepCount++
		}
	}

	fmt.Println("Number of steps to get to", endLoc, ":", stepCount)

	// ***** START HERE *****
	// Was trying recursion.. might not work due to step list??
	// node, count := followNode(firstNode, 1)

}

func followNode(n Node, d string, nm map[string]Node) Node {
	if d == "L" {
		fmt.Println("stepping LEFT from node", n.name, "to node", n.left)
		return nm[n.left]
	} else if d == "R" {
		fmt.Println("stepping RIGHT from node", n.name, "to node", n.right)
		return nm[n.right]
	}
	return Node{}
}

func part1() {

}

func part2() {

}

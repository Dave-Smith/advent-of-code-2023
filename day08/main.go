package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DesertMap struct {
	Moves     []string
	Nodes     map[string]Node
	currMove  int
	currNodes []string
}
type MapNode struct {
	Curr  string
	Left  string
	Right string
}
type Node struct {
	Left  string
	Right string
}

func NewDesertMap(dir []string, nodes []MapNode) DesertMap {
	n := make(map[string]Node)
	starts := make([]string, 0)
	for _, v := range nodes {
		n[v.Curr] = Node{v.Left, v.Right}
		if strings.HasSuffix(v.Curr, "A") {
			starts = append(starts, v.Curr)
		}
	}

	return DesertMap{dir, n, 0, starts}
}
func (m *DesertMap) NextNode() []string {
	next := m.Moves[m.currMove]
	if m.currMove+1 == len(m.Moves) {
		m.currMove = 0
		// fmt.Println("Finished a round of directions")
	} else {
		m.currMove++
	}
	//fmt.Print(next)
	nextNodes := make([]string, 0)
	for _, n := range m.currNodes {
		node := m.Nodes[n]
		if next == "L" {
			nextNodes = append(nextNodes, node.Left)
		} else {
			nextNodes = append(nextNodes, node.Right)
		}
	}

	m.currNodes = nextNodes
	return m.currNodes
	// for i := 0; i < len(m.Nodes); i++ {
	// 	if m.currNode == m.Nodes[i].Curr {
	// 		fmt.Print(m.currNode)
	// 		if next == "L" {
	// 			m.currNode = m.Nodes[i].Left
	// 			return m.currNode
	// 		}
	// 		m.currNode = m.Nodes[i].Right
	// 		return m.currNode
	// 	}
	// }
	// return ""
}

func main() {
	desert := readInput("./input.txt")

	res := 1
	for !isEndNode(desert.NextNode()) {
		res++
	}
	fmt.Println(res)
}

func isEndNode(nodes []string) bool {
	for i := 0; i < len(nodes); i++ {
		if !strings.HasSuffix(nodes[i], "Z") {
			return false
		}
	}
	return true
}

func readInput(filename string) DesertMap {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)

	lineScanner.Scan()
	header := lineScanner.Text()

	nodes := make([]MapNode, 0)

	replacer := strings.NewReplacer("=", "", "(", "", ",", "", ")", "")
	// remove := []string{"=", "(", ",", ")"}
	// replacer := func(r rune) bool {
	// 	for _, s := range remove {
	// 		if string(r) == s {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }

	for lineScanner.Scan() {
		line := lineScanner.Text()
		if len(line) == 0 {
			continue
		}

		line = replacer.Replace(line)
		fields := strings.Fields(line)
		nodes = append(nodes, MapNode{fields[0], fields[1], fields[2]})
	}
	return NewDesertMap(strings.Split(header, ""), nodes)
}

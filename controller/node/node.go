package node

import "fmt"

type Node struct {
	ID       string
	Name     string
	Parents  map[string]*Node
	Children map[string]*Node
}

func NewNode(id string, name string) *Node {
	node := Node{}
	node.ID = id
	node.Name = name
	node.Children = map[string]*Node{}
	node.Parents = map[string]*Node{}
	return &node
}
func (node *Node) DisplayNode() {
	fmt.Printf(" NODE ID :: %s , NODE NAME :: %s\n", node.ID, node.Name)
	fmt.Println("Parents ")
	for _, value := range node.Parents {
		fmt.Printf("%s --> %s\n", value.ID, node.ID)
	}
	fmt.Println("Children ")
	for _, value := range node.Children {
		fmt.Printf("%s --> %s\n", node.ID, value.ID)
	}
	fmt.Println("*********")
}

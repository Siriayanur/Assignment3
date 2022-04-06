package graph

import (
	"fmt"

	"github.com/Siriayanur/Assignment3/exceptions"
	"github.com/Siriayanur/Assignment3/model/node"
)

type Graph struct {
	nodes map[string]*node.Node
}

func CreateGraphInstance() *Graph {
	graph := Graph{}
	graph.nodes = make(map[string]*node.Node)
	return &graph
}
func (g *Graph) GetNode(parentID string) (*node.Node, error) {
	if g.nodes[parentID] == nil {
		return nil, exceptions.InvalidOperation("idNotExists", exceptions.ErrInvalidNode)
	}
	return g.nodes[parentID], nil
}

func (g *Graph) isCycleExists(childID string, parentID string) error {
	// get all ancestors
	parentAncestors, _ := g.GetAncestors(parentID)
	for _, currNode := range parentAncestors {
		if currNode.Id == childID {
			return exceptions.InvalidOperation("cycleExists", exceptions.ErrInvalidDependency)
		}
	}
	return nil
}

// Get parents
func (g *Graph) GetParents(nodeID string) ([]*node.Node, error) {
	currentNode, err := g.GetNode(nodeID)
	if err != nil {
		return nil, err
	}
	var parents []*node.Node
	// currentNode.Parents - <key,value> <ID,Node>
	for _, parentNode := range currentNode.Parents {
		parents = append(parents, parentNode)
	}
	return parents, nil
}

// Get children
func (g *Graph) GetChildren(nodeID string) ([]*node.Node, error) {
	currentNode, err := g.GetNode(nodeID)
	if err != nil {
		return nil, err
	}
	var children []*node.Node
	// currentNode.Children - map(<key,value>) ; map(<ID,Node>)
	for _, childNode := range currentNode.Children {
		children = append(children, childNode)
	}
	return children, nil
}

// Get ancestors.
func (g *Graph) getAncestorsHelper(currentNode *node.Node, visited map[string]bool, resStack *[]*node.Node) {
	if visited[currentNode.Id] {
		return
	}
	visited[currentNode.Id] = true
	*resStack = append(*resStack, currentNode)
	for _, parent := range currentNode.Parents {
		if !visited[parent.Id] {
			g.getAncestorsHelper(parent, visited, resStack)
		}
	}
}
func (g *Graph) GetAncestors(nodeID string) ([]*node.Node, error) {
	currentNode, err := g.GetNode(nodeID)
	if err != nil {
		return nil, err
	}
	parents := currentNode.Parents
	visited := make(map[string]bool)
	var resStack []*node.Node
	for _, parent := range parents {
		g.getAncestorsHelper(parent, visited, &resStack)
	}
	return resStack, nil
}

// Get descendents
func (g *Graph) getDescendentsHelper(currentNode *node.Node, visited map[string]bool, resStack *[]*node.Node) {
	if visited[currentNode.Id] {
		return
	}
	visited[currentNode.Id] = true
	*resStack = append(*resStack, currentNode)
	for _, child := range currentNode.Children {
		if !visited[child.Id] {
			g.getDescendentsHelper(child, visited, resStack)
		}
	}
}
func (g *Graph) GetDescendents(nodeID string) ([]*node.Node, error) {
	currentNode, err := g.GetNode(nodeID)
	if err != nil {
		return nil, err
	}
	children := currentNode.Children
	visited := make(map[string]bool)
	var resStack []*node.Node
	for _, child := range children {
		g.getDescendentsHelper(child, visited, &resStack)
	}
	return resStack, nil
}

// Add Node.
func (g *Graph) AddNode() error {
	var nodeID, nodeName string
	fmt.Println("Enter node id : ")
	fmt.Scanln(&nodeID)
	// check if the nodeId is unique
	if g.nodes[nodeID] != nil {
		return exceptions.InvalidOperation("idExists", exceptions.ErrInvalidNode)
	}
	fmt.Println("Enter node name : ")
	fmt.Scanln(&nodeName)
	newNode := node.NewNode(nodeID, nodeName)
	// add this newNode to the existing graph
	g.nodes[nodeID] = newNode
	fmt.Printf("Node %s added successfully\n", nodeID)
	// add dependencies if any
	fmt.Println("To add a dependency select y/yes")
	var choice string
	fmt.Scanln(&choice)
	if choice == "y" || choice == "yes" {
		err := g.AddDependency()
		if err != nil {
			return err
		}
	}
	return nil
}

// Add Dependency.
func (g *Graph) AddDependency() error {
	var parentID, childID string
	fmt.Println("Enter parent ID : ")
	fmt.Scanln(&parentID)
	fmt.Println("Enter child ID : ")
	fmt.Scanln(&childID)
	if g.nodes[parentID] == nil || g.nodes[childID] == nil {
		return exceptions.InvalidOperation("idNotExists", exceptions.ErrInvalidNode)
	}
	err := g.isCycleExists(childID, parentID)
	if err != nil {
		return err
	}
	childNode, err := g.GetNode(childID)
	if err != nil {
		return err
	}
	parentNode, err := g.GetNode(parentID)
	if err != nil {
		return err
	}
	childNode.Parents[parentID] = parentNode
	parentNode.Children[childID] = childNode
	return nil
}

// Delete Node.
func (g *Graph) DeleteNode(nodeID string) error {
	currentNode, err := g.GetNode(nodeID)
	if err != nil {
		return err
	}
	// get all parents
	parents := currentNode.Parents
	// get all children
	children := currentNode.Children
	// delete the concerned dependencies
	for _, parent := range parents {
		g.DeleteDependencyHelper(nodeID, parent.Id)
	}
	for _, child := range children {
		g.DeleteDependencyHelper(child.Id, nodeID)
	}
	// delete from graph
	delete(g.nodes, nodeID)
	return nil
}

// Delete dependency.
func (g *Graph) DeleteDependencyHelper(childID string, parentID string) error {
	if g.nodes[parentID] == nil || g.nodes[childID] == nil {
		return exceptions.InvalidOperation("idNotExists", exceptions.ErrInvalidNode)
	}
	childNode, _ := g.GetNode(childID)
	parentNode, _ := g.GetNode(parentID)
	delete(childNode.Parents, parentID)
	delete(parentNode.Children, childID)
	return nil
}
func (g *Graph) DeleteDependency() error {
	var parentID, childID string
	fmt.Println("Enter parent ID : ")
	fmt.Scanln(&parentID)
	fmt.Println("Enter child ID : ")
	fmt.Scanln(&childID)
	err := g.DeleteDependencyHelper(childID, parentID)
	return err
}

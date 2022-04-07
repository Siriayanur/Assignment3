package graph

import (
	"fmt"

	"github.com/Siriayanur/Assignment3/controller/node"
	"github.com/Siriayanur/Assignment3/exceptions"
)

type Graph struct {
	nodes map[string]*node.Node
}

// Get parents.
func (g *Graph) GetParents(nodeID string) ([]*node.Node, error) {
	currentNode, err := g.GetNode(nodeID)
	if err != nil {
		return nil, err
	}
	var parents []*node.Node = []*node.Node{}
	// currentNode.Parents - <key,value> <ID,Node>
	for _, parentNode := range currentNode.Parents {
		parents = append(parents, parentNode)
	}
	return parents, nil
}

// Get children.
func (g *Graph) GetChildren(nodeID string) ([]*node.Node, error) {
	currentNode, err := g.GetNode(nodeID)
	if err != nil {
		return nil, err
	}
	var children []*node.Node = []*node.Node{}
	// currentNode.Children - map(<key,value>) ; map(<ID,Node>)
	for _, childNode := range currentNode.Children {
		children = append(children, childNode)
	}
	return children, nil
}

// Get ancestors.
func (g *Graph) getAncestorsMain(currentNode *node.Node, visited map[string]bool, resStack *[]*node.Node) {
	if visited[currentNode.ID] {
		return
	}
	visited[currentNode.ID] = true
	*resStack = append(*resStack, currentNode)
	for _, parent := range currentNode.Parents {
		if !visited[parent.ID] {
			g.getAncestorsMain(parent, visited, resStack)
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
		g.getAncestorsMain(parent, visited, &resStack)
	}
	return resStack, nil
}

// Get descendents.
func (g *Graph) getDescendentsMain(currentNode *node.Node, visited map[string]bool, resStack *[]*node.Node) {
	if visited[currentNode.ID] {
		return
	}
	visited[currentNode.ID] = true
	*resStack = append(*resStack, currentNode)
	for _, child := range currentNode.Children {
		if !visited[child.ID] {
			g.getDescendentsMain(child, visited, resStack)
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
		g.getDescendentsMain(child, visited, &resStack)
	}
	return resStack, nil
}

// Add Node.
func (g *Graph) addNodeMain(nodeID string, nodeName string) error {
	// check if the nodeId is unique
	if g.nodes[nodeID] != nil {
		return exceptions.InvalidOperation("idExists", exceptions.ErrInvalidNode)
	}
	newNode := node.NewNode(nodeID, nodeName)
	// add this newNode to the existing graph
	g.nodes[nodeID] = newNode
	return nil
}
func (g *Graph) AddNode() error {
	var nodeID, nodeName string
	fmt.Println("Enter node id : ")
	fmt.Scanln(&nodeID)
	fmt.Println("Enter node name : ")
	fmt.Scanln(&nodeName)
	err := g.addNodeMain(nodeID, nodeName)
	if err != nil {
		return err
	}
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
func (g *Graph) addDependencyMain(parentID string, childID string) error {
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
func (g *Graph) AddDependency() error {
	var parentID, childID string
	fmt.Println("Enter parent ID : ")
	fmt.Scanln(&parentID)
	fmt.Println("Enter child ID : ")
	fmt.Scanln(&childID)
	exists, err := g.dependencyExists(parentID, childID)
	if err != nil {
		return err
	}
	if exists {
		fmt.Printf("Dependency already exists %s --> %s\n", parentID, childID)
		return nil
	}
	err = g.addDependencyMain(parentID, childID)
	if err != nil {
		return err
	}
	fmt.Printf("Dependency added successfully %s --> %s\n", parentID, childID)
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
		err := g.deleteDependencyMain(parent.ID, nodeID)
		if err != nil {
			return err
		}
	}
	for _, child := range children {
		err := g.deleteDependencyMain(nodeID, child.ID)
		if err != nil {
			return err
		}
	}
	// delete from graph
	delete(g.nodes, nodeID)
	fmt.Printf("Deleted node %s along with its dependencies\n", nodeID)
	return nil
}

// Delete dependency.
func (g *Graph) deleteDependencyMain(parentID string, childID string) error {
	if g.nodes[parentID] == nil || g.nodes[childID] == nil {
		return exceptions.InvalidOperation("idNotExists", exceptions.ErrInvalidNode)
	}
	childNode, _ := g.GetNode(childID)
	parentNode, _ := g.GetNode(parentID)
	// check if such a dependency exists
	if childNode.Parents[parentID] == nil {
		return exceptions.InvalidOperation("dependencyNotExists", exceptions.ErrInvalidDependency)
	}
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
	err := g.deleteDependencyMain(parentID, childID)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted dependency %s --> %s \n", parentID, childID)
	return nil
}

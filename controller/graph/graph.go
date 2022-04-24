package graph

import (
	"fmt"

	"github.com/Siriayanur/Assignment3/controller/node"
	"github.com/Siriayanur/Assignment3/exceptions"
)

type Graph struct {
	nodes map[string]*node.Node
}
type IGraph interface {
	GetParents(string) ([]*node.Node, error)
	GetChildren(string) ([]*node.Node, error)
	GetAncestors(string) ([]*node.Node, error)
	GetDescendents(string) ([]*node.Node, error)
	AddNodeHelper(string, string) error
	AddDependencyHelper(string, string) error
	DeleteNodeHelper(string) error
	DeleteDependencyHelper(string, string) error
}

// Get parents.
func (g *Graph) GetParents(nodeID string) ([]*node.Node, error) {
	currentNode, err := g.GetNode(nodeID)
	if err != nil {
		return nil, err
	}
	parents := make([]*node.Node, 0, len(currentNode.Parents))
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
	children := make([]*node.Node, 0, len(currentNode.Children))
	// currentNode.Children - map(<key,value>) ; map(<ID,Node>)
	for _, childNode := range currentNode.Children {
		children = append(children, childNode)
	}
	return children, nil
}

// Get ancestors.
func (g *Graph) GetAncestorsHelper(currentNode *node.Node, visited map[string]bool, ancestors *[]*node.Node) {
	if visited[currentNode.ID] {
		return
	}
	visited[currentNode.ID] = true
	*ancestors = append(*ancestors, currentNode)
	for _, parent := range currentNode.Parents {
		if !visited[parent.ID] {
			g.GetAncestorsHelper(parent, visited, ancestors)
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
	var ancestors []*node.Node
	for _, parent := range parents {
		g.GetAncestorsHelper(parent, visited, &ancestors)
	}
	return ancestors, nil
}

// Get descendents.
func (g *Graph) GetDescendentsHelper(currentNode *node.Node, visited map[string]bool, descendents *[]*node.Node) {
	if visited[currentNode.ID] {
		return
	}
	visited[currentNode.ID] = true
	*descendents = append(*descendents, currentNode)
	for _, child := range currentNode.Children {
		if !visited[child.ID] {
			g.GetDescendentsHelper(child, visited, descendents)
		}
	}
}
func (g *Graph) GetDescendents(nodeID string) ([]*node.Node, error) {
	currentNode, err := g.GetNode(nodeID)
	if err != nil {
		return nil, err
	}
	visited := make(map[string]bool)
	var descendents []*node.Node
	for _, child := range currentNode.Children {
		g.GetDescendentsHelper(child, visited, &descendents)
	}
	return descendents, nil
}

// Add Node.
func (g *Graph) AddNodeHelper(nodeID string, nodeName string) error {
	// check if the nodeId is unique
	if g.nodes[nodeID] != nil {
		return exceptions.InvalidOperation("idExists", exceptions.ErrInvalidNode)
	}
	newNode := node.NewNode(nodeID, nodeName)
	// add this newNode to the existing graph
	g.nodes[nodeID] = newNode
	return nil
}

// Add Dependency.
func (g *Graph) AddDependencyHelper(parentID string, childID string) error {
	// check if dependency already exists
	exists, err := g.dependencyExists(parentID, childID)
	if err != nil {
		return err
	}
	if exists {
		fmt.Printf("Dependency already exists %s --> %s\n", parentID, childID)
		return nil
	}
	if g.nodes[parentID] == nil || g.nodes[childID] == nil {
		return exceptions.InvalidOperation("idNotExists", exceptions.ErrInvalidNode)
	}
	err = g.isCycleExists(childID, parentID)
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
func (g *Graph) DeleteNodeHelper(nodeID string) error {
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
		err := g.DeleteDependencyHelper(parent.ID, nodeID)
		if err != nil {
			return err
		}
	}
	for _, child := range children {
		err := g.DeleteDependencyHelper(nodeID, child.ID)
		if err != nil {
			return err
		}
	}
	// delete from graph
	delete(g.nodes, nodeID)
	return nil
}

// Delete dependency.
func (g *Graph) DeleteDependencyHelper(parentID string, childID string) error {
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

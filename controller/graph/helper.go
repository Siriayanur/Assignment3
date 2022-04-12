package graph

import (
	"github.com/Siriayanur/Assignment3/controller/node"
	"github.com/Siriayanur/Assignment3/exceptions"
)

// used to hold all the node references.
func NewGraph() *Graph {
	graph := Graph{}
	graph.nodes = make(map[string]*node.Node)
	return &graph
}

// get node from ID.
func (g *Graph) GetNode(nodeID string) (*node.Node, error) {
	if g.nodes[nodeID] == nil {
		return nil, exceptions.InvalidOperation("idNotExists", exceptions.ErrInvalidNode)
	}
	return g.nodes[nodeID], nil
}

// check if a cycle exists between parent and child node before addition of dependency.
func (g *Graph) IsCycleExists(childID string, parentID string) error {
	// get all ancestors of parentNode
	parentAncestors, _ := g.GetAncestors(parentID)
	// if child node happens to be in the list, cycle exists
	for _, currNode := range parentAncestors {
		if currNode.ID == childID {
			return exceptions.InvalidOperation("cycleExists", exceptions.ErrInvalidDependency)
		}
	}
	return nil
}

// check if a dependency already exists between parentID and childID.
func (g *Graph) DependencyExists(parentID string, childID string) (bool, error) {
	childNode, err := g.GetNode(childID)
	if err != nil {
		return false, err
	}
	_, err = g.GetNode(parentID)
	if err != nil {
		return false, err
	}
	if childNode.Parents[parentID] != nil {
		return true, nil
	}
	return false, nil
}

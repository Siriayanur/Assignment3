package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestNodes() *Graph {
	g := CreateGraphInstance()
	_ = g.addNodeMain("1", "one")
	_ = g.addNodeMain("2", "two")
	_ = g.addNodeMain("3", "three")
	_ = g.addNodeMain("4", "four")
	_ = g.addNodeMain("5", "five")
	return g
}
func createTestGraph() *Graph {
	g := CreateGraphInstance()
	_ = g.addNodeMain("1", "one")
	_ = g.addNodeMain("2", "two")
	_ = g.addNodeMain("3", "three")
	_ = g.addNodeMain("4", "four")
	_ = g.addNodeMain("5", "five")
	_ = g.addNodeMain("6", "six")
	_ = g.addNodeMain("7", "seven")

	_ = g.addDependencyMain("1", "2")
	_ = g.addDependencyMain("1", "3")
	_ = g.addDependencyMain("2", "4")
	_ = g.addDependencyMain("2", "5")
	_ = g.addDependencyMain("5", "3")
	_ = g.addDependencyMain("5", "7")
	_ = g.addDependencyMain("3", "6")
	_ = g.addDependencyMain("7", "6")
	return g
}
func TestAddNode(t *testing.T) {
	g := CreateGraphInstance()
	err1 := g.addNodeMain("1", "one")
	err2 := g.addNodeMain("3@", "three")
	err3 := g.addNodeMain("1", "two")
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.NotNil(t, err3)
}
func TestAddDependency(t *testing.T) {
	g := createTestNodes()
	err1 := g.addDependencyMain("1", "2")
	err2 := g.addDependencyMain("1", "3")
	err3 := g.addDependencyMain("2", "4")
	err4 := g.addDependencyMain("4", "1")
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Nil(t, err3)
	require.NotNil(t, err4)
}
func TestGetParents(t *testing.T) {
	g := createTestGraph()
	// select a valid node in graph
	validNodeID := "6"
	// select an invalid node in graph
	invalidNodeID := "10"
	expParents := make(map[string]bool)
	// expected Parents for nodeIDPresent
	expParents["3"] = true
	expParents["7"] = true
	// get the parents for validNodeID
	parents, err1 := g.GetParents(validNodeID)
	require.Nil(t, err1)
	// error generated for invalid node
	_, err2 := g.GetParents(invalidNodeID)
	require.NotNil(t, err2)
	// parents should contain only the nodes that are in expParents
	require.Equal(t, len(parents), len(expParents))
	for _, node := range parents {
		require.Contains(t, expParents, node.ID)
	}
}
func TestGetChildren(t *testing.T) {
	g := createTestGraph()
	// select a valid node in graph
	validNodeID := "2"
	// not present in graph
	invalidNodeID := "10"
	// expected children for validNodeID
	expChildren := make(map[string]bool)
	expChildren["4"] = true
	expChildren["5"] = true
	children, err1 := g.GetChildren(validNodeID)
	require.Nil(t, err1)
	// error should be generated for invalidNodeID
	_, err2 := g.GetChildren(invalidNodeID)
	require.NotNil(t, err2)
	// children should contain only the nodes that are in expChildren
	require.Equal(t, len(children), len(expChildren))
	for _, child := range children {
		require.Contains(t, expChildren, child.ID)
	}
}
func TestGetAncestors(t *testing.T) {
	g := createTestGraph()
	// select a valid node in graph
	validNodeID := "7"
	// not present in graph
	invalidNodeID := "10"
	// expected children for validNodeID
	expAncestors := make(map[string]bool)
	expAncestors["5"] = true
	expAncestors["1"] = true
	expAncestors["2"] = true
	ancestors, err1 := g.GetAncestors(validNodeID)
	require.Nil(t, err1)
	// error should be generated for invalidNodeID
	_, err2 := g.GetAncestors(invalidNodeID)
	require.NotNil(t, err2)
	// children should contain only the nodes that are in expChildren
	require.Equal(t, len(ancestors), len(expAncestors))
	for _, ancestor := range ancestors {
		require.Contains(t, expAncestors, ancestor.ID)
	}
}
func TestGetDescendents(t *testing.T) {
	g := createTestGraph()
	// select a valid node in graph
	validNodeID := "5"
	// select an invalid node in graph
	invalidNodeID := "10"
	// expected descendents for validNodeID
	expDescendents := make(map[string]bool)
	expDescendents["3"] = true
	expDescendents["7"] = true
	expDescendents["6"] = true
	// error should be generated for invalidNodeID
	descendents, err1 := g.GetDescendents(validNodeID)
	require.Nil(t, err1)
	// throws error for the node that is not present
	_, err2 := g.GetDescendents(invalidNodeID)
	require.NotNil(t, err2)
	// children should contain only the nodes that are in expChildren
	require.Equal(t, len(descendents), len(expDescendents))
	for _, descendent := range descendents {
		require.Contains(t, expDescendents, descendent.ID)
	}
}
func TestDeleteDependency(t *testing.T) {
	g := createTestGraph()
	// select valid and invalid parent and child IDs
	parentID := "5"
	childID := "7"
	childIDNotPresent := "10"
	parentIDNotPresent := "9"
	// expected descendents for parentID
	expChildren := make(map[string]bool)
	expChildren["3"] = true
	// even if one of the IDs(parent/child) is not valid, err is thrown
	err1 := g.deleteDependencyMain(parentID, childID)
	err2 := g.deleteDependencyMain(parentIDNotPresent, childID)
	err3 := g.deleteDependencyMain(parentID, childIDNotPresent)
	// error should not be nil for parentIDNotPresent and childIDNotPresent
	require.NotNil(t, err2)
	require.NotNil(t, err3)
	require.Nil(t, err1)
	// get descendents after removing dependency between valid parent --> valid child
	children, err := g.GetChildren(parentID)
	require.Nil(t, err)
	require.Equal(t, len(children), len(expChildren))
	for _, child := range children {
		require.Contains(t, expChildren, child.ID)
	}
}
func TestDeleteNode(t *testing.T) {
	g := createTestGraph()
	// select a valid node from graph to delete
	nodePresentID := "7"
	// select an invalid node from graph
	nodeNotPresentID := "10"
	// select a valid node whose descendents included nodePresentID
	validNodeID := "5"
	// expected descendents for validNodeID after deleteNode() action
	expDescendents := make(map[string]bool)
	expDescendents["3"] = true
	expDescendents["6"] = true
	// perform deleteNode action on both valid and invalid nodes
	err1 := g.DeleteNode(nodePresentID)
	err2 := g.DeleteNode(nodeNotPresentID)
	// error should not be nil for nodeNotPresentID
	require.Nil(t, err1)
	require.NotNil(t, err2)
	// get descendents of validNode after deleting nodePresentID
	descendents, err := g.GetDescendents(validNodeID)
	require.Nil(t, err)
	// check if length and each element of descendents matched with expected descendents
	require.Equal(t, len(descendents), len(expDescendents))
	for _, descendent := range descendents {
		require.Contains(t, expDescendents, descendent.ID)
	}
}

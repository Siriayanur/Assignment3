package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestNodes() *Graph {
	g := NewGraph()
	_ = g.AddNodeHelper("1", "one")
	_ = g.AddNodeHelper("2", "two")
	_ = g.AddNodeHelper("3", "three")
	_ = g.AddNodeHelper("4", "four")
	_ = g.AddNodeHelper("5", "five")
	return g
}
func createTestGraph() *Graph {
	g := NewGraph()
	_ = g.AddNodeHelper("1", "one")
	_ = g.AddNodeHelper("2", "two")
	_ = g.AddNodeHelper("3", "three")
	_ = g.AddNodeHelper("4", "four")
	_ = g.AddNodeHelper("5", "five")
	_ = g.AddNodeHelper("6", "six")
	_ = g.AddNodeHelper("7", "seven")

	_ = g.AddDependencyHelper("1", "2")
	_ = g.AddDependencyHelper("1", "3")
	_ = g.AddDependencyHelper("2", "4")
	_ = g.AddDependencyHelper("2", "5")
	_ = g.AddDependencyHelper("5", "3")
	_ = g.AddDependencyHelper("5", "7")
	_ = g.AddDependencyHelper("3", "6")
	_ = g.AddDependencyHelper("7", "6")
	return g
}
func TestAddNode(t *testing.T) {
	g := NewGraph()
	err1 := g.AddNodeHelper("1", "one")
	err2 := g.AddNodeHelper("3@", "three")
	err3 := g.AddNodeHelper("1", "two")
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.NotNil(t, err3)
}
func TestAddDependencyValid(t *testing.T) {
	g := createTestNodes()
	err1 := g.AddDependencyHelper("1", "2")
	err2 := g.AddDependencyHelper("1", "3")
	err3 := g.AddDependencyHelper("2", "4")
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Nil(t, err3)
}
func TestAddDependencyInvalid(t *testing.T) {
	g := createTestNodes()
	err1 := g.AddDependencyHelper("1", "2")
	err2 := g.AddDependencyHelper("1", "3")
	err3 := g.AddDependencyHelper("2", "4")
	err4 := g.AddDependencyHelper("4", "1")
	require.Nil(t, err1)
	require.Nil(t, err2)
	require.Nil(t, err3)
	require.NotNil(t, err4)
}
func TestGetParentsValid(t *testing.T) {
	g := createTestGraph()
	// select a valid node in graph
	validNodeID := "6"
	expParents := make(map[string]bool)
	// expected Parents for nodeIDPresent
	expParents["3"] = true
	expParents["7"] = true
	// get the parents for validNodeID
	parents, err1 := g.GetParents(validNodeID)
	require.Nil(t, err1)
	// parents should contain only the nodes that are in expParents
	require.Equal(t, len(parents), len(expParents))
	for _, node := range parents {
		require.Contains(t, expParents, node.ID)
	}
}
func TestGetParentsInvalid(t *testing.T) {
	g := createTestGraph()
	// select an invalid node in graph
	invalidNodeID := "10"
	// error generated for invalid node
	_, err2 := g.GetParents(invalidNodeID)
	require.NotNil(t, err2)
}
func TestGetChildrenValid(t *testing.T) {
	g := createTestGraph()
	// select a valid node in graph
	validNodeID := "2"
	// expected children for validNodeID
	expChildren := make(map[string]bool)
	expChildren["4"] = true
	expChildren["5"] = true
	children, err1 := g.GetChildren(validNodeID)
	require.Nil(t, err1)
	// children should contain only the nodes that are in expChildren
	require.Equal(t, len(children), len(expChildren))
	for _, child := range children {
		require.Contains(t, expChildren, child.ID)
	}
}
func TestGetChildrenInvalid(t *testing.T) {
	g := createTestGraph()
	// not present in graph
	invalidNodeID := "10"
	// error should be generated for invalidNodeID
	_, err2 := g.GetChildren(invalidNodeID)
	require.NotNil(t, err2)
}
func TestGetAncestorsValid(t *testing.T) {
	g := createTestGraph()
	// select a valid node in graph
	validNodeID := "7"
	// expected children for validNodeID
	expAncestors := make(map[string]bool)
	expAncestors["5"] = true
	expAncestors["1"] = true
	expAncestors["2"] = true
	ancestors, err1 := g.GetAncestors(validNodeID)
	require.Nil(t, err1)
	// children should contain only the nodes that are in expChildren
	require.Equal(t, len(ancestors), len(expAncestors))
	for _, ancestor := range ancestors {
		require.Contains(t, expAncestors, ancestor.ID)
	}
}
func TestGetAncestorsInvalid(t *testing.T) {
	g := createTestGraph()
	// not present in graph
	invalidNodeID := "10"
	// error should be generated for invalidNodeID
	_, err2 := g.GetAncestors(invalidNodeID)
	require.NotNil(t, err2)
}
func TestGetDescendentsValid(t *testing.T) {
	g := createTestGraph()
	// select a valid node in graph
	validNodeID := "5"
	// expected descendents for validNodeID
	expDescendents := make(map[string]bool)
	expDescendents["3"] = true
	expDescendents["7"] = true
	expDescendents["6"] = true
	// error should be generated for invalidNodeID
	descendents, err1 := g.GetDescendents(validNodeID)
	require.Nil(t, err1)
	// children should contain only the nodes that are in expChildren
	require.Equal(t, len(descendents), len(expDescendents))
	for _, descendent := range descendents {
		require.Contains(t, expDescendents, descendent.ID)
	}
}
func TestGetDescendentsInvalid(t *testing.T) {
	g := createTestGraph()
	// select an invalid node in graph
	invalidNodeID := "10"
	// throws error for the node that is not present
	_, err2 := g.GetDescendents(invalidNodeID)
	require.NotNil(t, err2)
}
func TestDeleteDependencyValid(t *testing.T) {
	g := createTestGraph()
	// select valid and invalid parent and child IDs
	parentID := "5"
	childID := "7"
	// expected descendents for parentID
	expChildren := make(map[string]bool)
	expChildren["3"] = true
	// even if one of the IDs(parent/child) is not valid, err is thrown
	err1 := g.DeleteDependencyHelper(parentID, childID)
	require.Nil(t, err1)
	// get descendents after removing dependency between valid parent --> valid child
	children, err := g.GetChildren(parentID)
	require.Nil(t, err)
	require.Equal(t, len(children), len(expChildren))
	for _, child := range children {
		require.Contains(t, expChildren, child.ID)
	}
}
func TestDeleteDependencyInvalid(t *testing.T) {
	g := createTestGraph()
	// select valid and invalid parent and child IDs
	parentID := "5"
	childID := "7"
	childIDNotPresent := "10"
	parentIDNotPresent := "9"
	// even if one of the IDs(parent/child) is not valid, err is thrown
	err1 := g.DeleteDependencyHelper(parentIDNotPresent, childID)
	err2 := g.DeleteDependencyHelper(parentID, childIDNotPresent)
	// error should not be nil for parentIDNotPresent and childIDNotPresent
	require.NotNil(t, err1)
	require.NotNil(t, err2)
}
func TestDeleteNodeValid(t *testing.T) {
	g := createTestGraph()
	// select a valid node from graph to delete
	nodePresentID := "7"
	// select a valid node whose descendents included nodePresentID
	validNodeID := "5"
	// expected descendents for validNodeID after deleteNode() action
	expDescendents := make(map[string]bool)
	expDescendents["3"] = true
	expDescendents["6"] = true
	// perform deleteNode action on both valid and invalid nodes
	err1 := g.DeleteNodeHelper(nodePresentID)
	// error should not be nil for nodeNotPresentID
	require.Nil(t, err1)
	// get descendents of validNode after deleting nodePresentID
	descendents, err := g.GetDescendents(validNodeID)
	require.Nil(t, err)
	// check if length and each element of descendents matched with expected descendents
	require.Equal(t, len(descendents), len(expDescendents))
	for _, descendent := range descendents {
		require.Contains(t, expDescendents, descendent.ID)
	}
}
func TestDeleteNodeInvalid(t *testing.T) {
	g := createTestGraph()
	// select an invalid node from graph
	nodeNotPresentID := "10"
	// perform deleteNode action on both valid and invalid nodes
	err2 := g.DeleteNodeHelper(nodeNotPresentID)
	// error should not be nil for nodeNotPresentID
	require.NotNil(t, err2)
}

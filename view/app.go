package view

import (
	"fmt"
	"os"

	"github.com/Siriayanur/Assignment3/controller/graph"
	"github.com/Siriayanur/Assignment3/exceptions"
)

func RunApp() {
	var choice int
	exceptions.CreateErrorStatements()
	// create graph instance.
	g := graph.CreateGraphInstance()
	for {
		displayMenu()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			nodeID := getNodeID()
			parents, err := g.GetParents(nodeID)
			checkError(err)
			displayNodes(parents, nodeID, "PARENTS")
		case 2:
			nodeID := getNodeID()
			children, err := g.GetChildren(nodeID)
			checkError(err)
			displayNodes(children, nodeID, "CHILDREN")
		case 3:
			nodeID := getNodeID()
			ancestors, err := g.GetAncestors(nodeID)
			checkError(err)
			displayNodes(ancestors, nodeID, "ANCESTORS")
		case 4:
			nodeID := getNodeID()
			descendents, err := g.GetDescendents(nodeID)
			checkError(err)
			displayNodes(descendents, nodeID, "DESCENDENTS")
		case 5:
			err := g.DeleteDependency()
			checkError(err)
		case 6:
			err := g.DeleteNode(getNodeID())
			checkError(err)
		case 7:
			err := g.AddDependency()
			checkError(err)
		case 8:
			err := g.AddNode()
			checkError(err)
		default:
			fmt.Println("Invalid choice")
			os.Exit(1)
		}
	}
}

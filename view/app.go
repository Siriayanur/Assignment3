package view

import (
	"fmt"
	"os"

	"github.com/Siriayanur/Assignment3/controller/graph"
	"github.com/Siriayanur/Assignment3/exceptions"
)

type App struct {
	data graph.IGraph
}

func (app *App) DeleteDependency() error {
	var parentID, childID string
	fmt.Println("Enter parent ID : ")
	fmt.Scanln(&parentID)
	fmt.Println("Enter child ID : ")
	fmt.Scanln(&childID)
	err := app.data.DeleteDependencyHelper(parentID, childID)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted dependency %s --> %s \n", parentID, childID)
	return nil
}
func (app *App) AddDependency() error {
	var parentID, childID string
	fmt.Println("Enter parent ID : ")
	fmt.Scanln(&parentID)
	fmt.Println("Enter child ID : ")
	fmt.Scanln(&childID)
	err := app.data.AddDependencyHelper(parentID, childID)
	if err != nil {
		return err
	}
	fmt.Printf("Dependency added successfully %s --> %s\n", parentID, childID)
	return nil
}
func (app *App) AddNode() error {
	var nodeID, nodeName string
	fmt.Println("Enter node id : ")
	fmt.Scanln(&nodeID)
	fmt.Println("Enter node name : ")
	fmt.Scanln(&nodeName)
	err := app.data.AddNodeHelper(nodeID, nodeName)
	if err != nil {
		return err
	}
	fmt.Printf("Node %s added successfully\n", nodeID)
	// add dependencies if any
	fmt.Println("To add a dependency select y/yes")
	var choice string
	fmt.Scanln(&choice)
	if choice == "y" || choice == "yes" {
		err := app.AddDependency()
		if err != nil {
			return err
		}
	}
	return nil
}
func (app *App) DeleteNode(nodeID string) error {
	err := app.data.DeleteNodeHelper(nodeID)
	if err != nil {
		return err
	}
	fmt.Printf("Deleted node %s along with its dependencies\n", nodeID)
	return nil
}
func RunApp() {
	var choice int
	exceptions.CreateErrorStatements()
	// create graph instance.
	g := graph.NewGraph()
	app := App{g}
	for {
		displayMenu()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			nodeID := getNodeID()
			parents, err := app.data.GetParents(nodeID)
			checkError(err)
			displayNodes(parents, nodeID, "PARENTS")
		case 2:
			nodeID := getNodeID()
			children, err := app.data.GetChildren(nodeID)
			checkError(err)
			displayNodes(children, nodeID, "CHILDREN")
		case 3:
			nodeID := getNodeID()
			ancestors, err := app.data.GetAncestors(nodeID)
			checkError(err)
			displayNodes(ancestors, nodeID, "ANCESTORS")
		case 4:
			nodeID := getNodeID()
			descendents, err := app.data.GetDescendents(nodeID)
			checkError(err)
			displayNodes(descendents, nodeID, "DESCENDENTS")
		case 5:
			err := app.DeleteDependency()
			checkError(err)
		case 6:
			err := app.DeleteNode(getNodeID())
			checkError(err)
		case 7:
			err := app.AddDependency()
			checkError(err)
		case 8:
			err := app.AddNode()
			checkError(err)
		default:
			fmt.Println("Invalid choice")
			os.Exit(1)
		}
	}
}

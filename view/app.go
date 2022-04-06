package view

import (
	"fmt"

	"github.com/Siriayanur/Assignment3/exceptions"
	"github.com/Siriayanur/Assignment3/model/graph"
)

func RunApp() {
	var choice int
	exceptions.CreateErrorStatements()
	g := graph.CreateGraphInstance()
	// create graph
	for {
		displayMenu()
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var id string
			fmt.Println("Enter ID of node : ")
			fmt.Scanln(&id)
			parents, err := g.GetParents(id)
			if err != nil {
				fmt.Println(err)
				// os.Exit(1)
			}
			displayNodes(parents)
		case 2:
			var id string
			fmt.Println("Enter ID of node : ")
			fmt.Scanln(&id)
			children, err := g.GetChildren(id)
			if err != nil {
				fmt.Println(err)
				// os.Exit(1)
			}
			displayNodes(children)
		case 3:
			var id string
			fmt.Println("Enter ID of node : ")
			fmt.Scanln(&id)
			ancestors, err := g.GetAncestors(id)
			if err != nil {
				fmt.Println(err)
				// os.Exit(1)
			}
			displayNodes(ancestors)
		case 4:
			var id string
			fmt.Println("Enter ID of node : ")
			fmt.Scanln(&id)
			descendents, err := g.GetDescendents(id)
			if err != nil {
				fmt.Println(err)
				// os.Exit(1)
			}
			displayNodes(descendents)
		case 5:
			err := g.DeleteDependency()
			if err != nil {
				fmt.Println(err)
				// os.Exit(1)
			}
		case 6:
			var id string
			fmt.Println("Enter ID of node : ")
			fmt.Scanln(&id)
			err := g.DeleteNode(id)
			if err != nil {
				fmt.Println(err)
				// os.Exit(1)
			}
		case 7:
			err := g.AddDependency()
			if err != nil {
				fmt.Println(err)
				// os.Exit(1)
			}
		case 8:
			err := g.AddNode()
			if err != nil {
				fmt.Println(err)
				// os.Exit(1)
			}
		case 9:
			fmt.Println("Invalid choice")
			// os.Exit(1)
		}
	}
}

package view

import (
	"fmt"

	"github.com/Siriayanur/Assignment3/model/node"
)

func displayMenu() {
	fmt.Println("| MENU |")
	fmt.Println("1. GET IMMEDIATE PARENTS |")
	fmt.Println("2. GET IMMEDIATE CHILDREN |")
	fmt.Println("3. GET THE ANCESTORS |")
	fmt.Println("4. GET THE DESCENDENTS |")
	fmt.Println("5. DELETE DEPENDENCY OF A NODE | ")
	fmt.Println("6. DELETE NODE FROM TREE |")
	fmt.Println("7. ADD NEW DEPENDENCY |")
	fmt.Println("8. ADD NEW NODE |")
}
func displayNodes(nodes []*node.Node) {
	fmt.Println("Node ID | Node Name")
	for _, currNode := range nodes {
		fmt.Printf(" %s | %s \n", currNode.Id, currNode.Name)
		// currNode.DisplayNode()
	}

}

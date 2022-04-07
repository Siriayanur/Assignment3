package view

import (
	"fmt"
	"os"

	"github.com/Siriayanur/Assignment3/controller/node"
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
	fmt.Println("")
}
func displayNodes(nodes []*node.Node, nodeID string, relation string) {
	if len(nodes) == 0 {
		fmt.Printf("No %s of %s\n\n", relation, nodeID)
		return
	}
	fmt.Printf("%s of %s\n\n", relation, nodeID)
	fmt.Println("NODE ID\tNODE NAME")
	for _, currNode := range nodes {
		fmt.Printf(" %s | %s \n", currNode.ID, currNode.Name)
	}
	fmt.Println()
}
func getNodeID() string {
	var id string
	fmt.Println("Enter ID of node : ")
	fmt.Scanln(&id)
	return id
}
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

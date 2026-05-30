package main

import "fmt"

func main() {
	activeNodes := []string{"Node A", "Node B"}
	
	activeNodes = append(activeNodes, "Node C")
	
	fmt.Printf("Node: %s\n", activeNodes[1])
	fmt.Printf("Active Nodes: %d\n", len(activeNodes))
}

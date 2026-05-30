package main

import (
	"fmt"
)

type Node struct {
Name string
IP string
isActive bool
}

func main() {
	n1 := Node{
	Name: "Node-Alpha",
	IP: "192.168.1.1",
	isActive: true,
	}
	
	n1.isActive = false
	
	if n1.isActive {
		fmt.Printf("Node: %s | Status: %t on IP: %s\n", n1.Name, n1.isActive, n1.IP)
	} else {
		fmt.Printf("Node: %s | Status: %t\n", n1.Name, n1.isActive)
	}

}

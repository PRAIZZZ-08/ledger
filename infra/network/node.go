package network // Label the box

import "fmt"

type Node struct { // Capital 'N' makes it Exported (Public)
    ID int         // Capital 'I' makes it Exported
    IP string      // Capital 'I' makes it Exported
}

// Moved from main. This belongs to the Node in this package.
func (n Node) GetConnectionInfo() string {
    return fmt.Sprintf("Node %d reachable at %s", n.ID, n.IP)
}

func (n Node) Report() string {
    return n.GetConnectionInfo() // Reuse internal logic
}

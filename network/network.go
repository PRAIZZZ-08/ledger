package network

type Node struct {
 ID int
 IP string
}

func (n Node) Report() string {
return fmt.Sprintf("Node %d reachable at %s\n", n.ID, n.IP)
}


package main

import  "fmt"

type Reportable interface {
 Report() string 
}

type Node struct {
 ID int
 IP string
}

type Entry struct {
 ID    int
 Value float64
}

func (e Entry) Report() string {
 return fmt.Sprintf("Entry %d has value $[%.2f]", e.ID, e.Value)   
}

func (n Node) Report() string {
return fmt.Sprintf("Node %d reachable at %s\n", n.ID, n.IP)
}

func LogComponent(arg Reportable) {
fmt.Println(arg.Report())
}

func main() {
e1 := Entry{ID: 46, Value: 65}
n1 := Node{ID: 234, IP: "192.168.0.1"}

LogComponent(e1)
LogComponent(n1)
}

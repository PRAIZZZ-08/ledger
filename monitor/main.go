package main

import  ( 
	"fmt"
	"monitor/network"
	"time"
	
	//"github.com/fatih/color"
)

type Reportable interface {
 Report() string 
}

type Entry struct {
 ID    int
 Value float64
}

func (e Entry) Report() string {
 return fmt.Sprintf("Entry %d has value $[%.2f]", e.ID, e.Value)   
}


func LogComponent(arg Reportable) {
rawReport := arg.Report()
coloredReport := color.GreenString(rawReport)
time.Sleep(1 * time.Second)
fmt.Println(coloredReport)
}

func main() {
e1 := Entry{ID: 46, Value: 65}
n1 := network.Node{ID: 234, IP: "192.168.0.1"}

go LogComponent(e1)
go LogComponent(n1)
time.Sleep(2 * time.Second)
}

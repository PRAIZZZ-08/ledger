package main

import  ( 
	"fmt"
	"monitor/network"
	"time"
	"sync"
	
	"github.com/fatih/color"
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


func LogComponent(arg Reportable, res chan string) {
    var id int
    switch v := arg.(type) {
    case Entry:
        id = v.ID
    case network.Node:
        id = v.ID
    }
    if id == 234 {
        time.Sleep(3 * time.Second)
    }
    
    rawReport := arg.Report()
    coloredReport := color.GreenString(rawReport)
    
    fmt.Println(coloredReport)   
    res <- rawReport
}

func main() {
var wg sync.WaitGroup
resChan :=  make(chan string)

e1 := []Entry{
	{ID: 46, Value: 65.1},
	{ID: 34, Value: 93.4},
	{ID: 96, Value: 86.5},
	{ID: 234, Value: 76.0},
}

n1 := network.Node{ID: 234, IP: "192.168.0.1"}

for _, p := range e1 {
        go LogComponent(p, resChan)
        go LogComponent(n1, resChan)
   }

  for i := 0; i < len(e1); i++ {
   select {
    case msg := <-resChan:
     fmt.Println("Received:", msg)
     case <-time.After(1 * time.Second):
     fmt.Println("ERROR: Component response timed out.")
   }
  }
}

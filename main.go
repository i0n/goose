package main

import(
  "fmt"
  //"os"
  //"github.com/i0n/goose/lib/BlueDragonX/go-supervisor/supervisor"
) 

//func main() {
//    done := make(chan bool)
//    events := make(chan supervisor.Event)
//    evl := supervisor.NewListener(os.Stdin, os.Stdout)
//
//    go func() {
//        for event := range events {
//            fmt.Fprintf(os.Stderr, "Got event: %s\n", event)
//        }
//        done <- true
//    }()
//
//    evl.Run(events)
//    close(events)
//    <-done
//}

func main() {
    fmt.Println("Hello, 世界")
}

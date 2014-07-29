package main

import(
  "fmt"
  "os"
  "syscall"
  "github.com/i0n/goose/lib/BlueDragonX/go-supervisor/supervisor"
)

//func main() {
//  fmt.Fprintf(os.Stderr, "Killing!\n")
//  syscall.Kill(1, syscall.SIGKILL)
//}

func main() {
    done := make(chan bool)
    events := make(chan supervisor.Event)
    evl := supervisor.NewListener(os.Stdin, os.Stdout)

    go func() {
        for event := range events {
            fmt.Fprintf(os.Stderr, "Got event: %s\n", event)
            if event.Name() == "PROCESS_STATE_FATAL" {
              fmt.Fprintf(os.Stderr, "Killing!\n")
              // Trying to call -1 here as you cannot stop pid 1 by calling '1'
              syscall.Kill(-1, syscall.SIGKILL)
              //syscall.Kill(syscall.Getpid(), syscall.SIGKILL)
            } else {
              fmt.Fprintf(os.Stderr, "Non Fatal!\n")
            }
        }
        done <- true
    }()

    evl.Run(events)
    close(events)
    <-done
}

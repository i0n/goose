package main

import(
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "time"
  "github.com/i0n/goose/lib/BlueDragonX/go-supervisor/supervisor"
)

func main() {
    done := make(chan bool)
    events := make(chan supervisor.Event)
    evl := supervisor.NewListener(os.Stdin, os.Stdout)
    sigs := make(chan os.Signal, 1)

    signal.Notify(sigs, syscall.SIGTERM)

    go func() {
        fmt.Fprintf(os.Stderr, "SIGTERM received\n")
        time.Sleep(3 * time.Second)
        syscall.Kill(-1, syscall.SIGKILL)
        done <- true
    }()

    go func() {
      for event := range events {
        fmt.Fprintf(os.Stderr, "Got event: %s\n", event)
          if event.Name() == "PROCESS_STATE_FATAL" {
            fmt.Fprintf(os.Stderr, "Killing!\n")
              // Calling -1 (stop everything) here as you cannot stop pid 1 by calling '1'
              syscall.Kill(-1, syscall.SIGTERM)
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

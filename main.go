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
  signals := make(chan os.Signal, 1)
  signal.Notify(signals, syscall.SIGTERM)
  args := os.Args[1:]

  for _, arg := range args {
    if arg == "-v" || arg == "--version" {
      fmt.Fprintf(os.Stdout, "%s\n", Version)
      return
    }
  }

  go func() {
    for signal := range signals {
      _ = signal
      fmt.Fprintf(os.Stderr, "SIGTERM received\n")
      time.Sleep(3 * time.Second)
      syscall.Kill(-1, syscall.SIGKILL)
    }
    done <- true
  }()

  go func() {
    for event := range events {
      if event.Name() == "PROCESS_STATE_FATAL" {
        // Calling -1 (stop everything) here as you cannot stop pid 1 by calling '1'
        syscall.Kill(-1, syscall.SIGTERM)
      }
    }
    done <- true
  }()

  evl.Run(events)
  close(events)
  <-done
}

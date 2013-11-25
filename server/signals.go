package server

import (
  "os"
  "fmt"
  "log"
  "time"
  "strconv"
  "syscall"
  "os/exec"
  "os/signal"
  "beatstorm/core"
)

func setupSignals() {
  setupRestartSignal()
  setupPID()
}

func setupRestartSignal() {
  sig := make(chan os.Signal, 1)
  signal.Notify(sig, syscall.SIGINT)
  go func() {
    <-sig
    if core.GetConfig().Env == core.PRODUCTION {
      root, _ := os.Getwd()
      cmd := exec.Command(root + "/start.sh")
      if err := cmd.Run(); err != nil { fmt.Println("Starting new process", err) }
      time.Sleep(30 * time.Second)
    }
    os.Exit(0)
  }()
}

func setupPID() {
  root, _ := os.Getwd()
  fpid, err := os.Create(root + "/beatstorm.pid")
  if err != nil { log.Println("failed to open beatstorm.pid") }
  fpid.Write([]byte(strconv.Itoa(os.Getpid())))
  fpid.Close()
}

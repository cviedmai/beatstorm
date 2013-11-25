package server

import (
  "os"
  "log"
  "time"
  "bytes"
  "strconv"
  "runtime"
)

func startStats(filename string) {
  go snapshot(filename)
}

func snapshot(filename string) {
  buffer := new(bytes.Buffer)
  for {
    time.Sleep(time.Minute)
    buffer.Reset()
    buffer.WriteString(`{"goroutines":` + strconv.Itoa(runtime.NumGoroutine()) + `}`)
    file, e := os.Create(filename)
    if e != nil {
      log.Println("Could not write stats: ", e)
    } else {
      file.Write(buffer.Bytes())
      file.Close()
    }
  }
}

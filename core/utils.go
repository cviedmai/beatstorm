package core

import (
  "os"
)

func GetPath(filename string) string {
  rootPath := os.Getenv("BS_ROOT")
  if len(rootPath) != 0 {
    return rootPath + "/" + filename
  }
  return filename
}

func GetDir(dirname string) string {
  rootPath := os.Getenv("BS_ROOT")
  if len(rootPath) != 0 {
    return rootPath + "/" + dirname + "/"
  }
  return dirname + "/"
}

package tests

import (
  "log"
  "beatstorm/core"
  "github.com/garyburd/redigo/redis"
)

var (
  conn redis.Conn
)

func init() {
  var err error
  c := core.GetConfig().Redis
  conn, err = redis.Dial("tcp", c.Host)
  if err != nil { log.Println(err) }
  if c.Db != 0 {
    if _, err = conn.Do("select", c.Db); err != nil { log.Println(err) }
  }
  conn.Do("flushdb")
}

func Redis() redis.Conn {
  return conn
}

package redis

import (
  "log"
  "beatstorm/core"
  "github.com/garyburd/redigo/redis"
)

var (
  Db *redis.Pool
)

func init() {
  c := core.GetConfig().Redis
  Db = &redis.Pool{
    MaxIdle: c.PoolSize,
    Dial: func() (redis.Conn, error) {
      conn, err := redis.Dial("tcp", c.Host)
      if err != nil {
        log.Println("Redis could not connect: ", err)
        return nil, err
      }
      if c.Db != 0 {
        if _, err = conn.Do("select", c.Db); err != nil {
          log.Println("Redis could not select db: ", err)
          return nil, err
        }
      }
      return conn, nil
    },
  }

  conn := Db.Get()
  defer conn.Close()
  if r, e := conn.Do("Ping"); r != "PONG" {
    log.Fatal("Redis connection failed: ", e)
  }
}

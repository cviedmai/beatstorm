package tests

import (
  "beatstorm/core/cache"
  "beatstorm/core/redis"
  "beatstorm/core/postgres"
)

func DbCleanUp(tables ...string) {
  cache.Clear()

  r := redis.Db.Get()
  defer r.Close()
  r.Do("flushdb")

  for _, t := range tables {
    postgres.Db.Exec("truncate table " + t)
  }
}

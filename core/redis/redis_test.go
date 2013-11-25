package redis

import (
  "testing"
  "github.com/cviedmai/gspec"
  "github.com/garyburd/redigo/redis"
)

func TestRedisWorks(t *testing.T) {
  spec := gspec.New(t)
  r := Db.Get()
  defer r.Close()
  r.Do("set", "name", "tyler")
  result, _ := redis.String(r.Do("get", "name"))
  spec.Expect(result).ToEqual("tyler")
}

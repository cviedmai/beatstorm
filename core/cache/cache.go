package cache

import (
  "time"
  "github.com/cviedmai/ccache"
)

var cache = ccache.New(ccache.Configure().Size(1024 * 1024 * 1024))

func Get(key string) interface{} {
  return cache.Get(key)
}

func Set(key string, value interface{}, duration time.Duration) {
  cache.Set(key, value, duration)
}

func Delete(key string) {
  cache.Delete(key)
}

func Fetch(key string, duration time.Duration, miss func() (interface{}, error)) (interface{}, error) {
  return cache.Fetch(key, duration, miss)
}

func FetchString(key string, duration time.Duration, miss func() (string, error)) (string, error) {
  item := cache.Get(key)
  if item != nil { return item.(string), nil }
  value, err := miss()
  if err != nil {
    cache.Set(key, value, duration)
  }
  return value, err
}

func Clear() {
  cache.Clear()
}

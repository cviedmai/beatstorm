package beatvector

import (
  "time"
  "testing"
  "beatstorm/tests"
  "beatstorm/core/cache"
  "github.com/cviedmai/gspec"
)

func TestJsonFromCache(t *testing.T) {
  tests.DbCleanUp("beatvectors")
  cache.Set("bv:nice_uri:json", "nice_json", time.Minute)
  bv := New()
  bv.Uri = "nice_uri"
  gspec.New(t).Expect(bv.Json()).ToEqual("nice_json")
}

func TestJsonFromBeatVector(t *testing.T) {
  tests.DbCleanUp("beatvectors")
  bv := New()
  bv.Uri = "nice_uri"
  bv.Data = "nice_data"
  gspec.New(t).Expect(bv.Json()).ToEqual(`{"uri":"nice_uri","data":"nice_data"}`)
}

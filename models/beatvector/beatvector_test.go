package beatvector

import (
  "time"
  "testing"
  "beatstorm/tests"
  "beatstorm/core/cache"
  "github.com/cviedmai/gspec"
)

func TestFindReturnsNilIfNotFound(t *testing.T) {
  tests.DbCleanUp("beatvectors")
  x := Find("caracol")
  gspec.New(t).Expect(x).ToBeNil()
}

func TestFindsBeatVectorFromCache(t *testing.T) {
  tests.DbCleanUp("beatvectors")
  bv := Builder().Uri("nice_uri").Data("nice_data").Bv
  cache.Set("bv:nice_uri", bv, time.Minute)
  gspec.New(t).Expect(Find("nice_uri")).ToEqual(bv)
}

func TestFindsBeatVectorFromDb(t *testing.T) {
  spec := gspec.New(t)
  tests.DbCleanUp("beatvectors")
  insertFakeBeatVector()
  bv := Find("nice_uri")
  spec.Expect(bv.Uri).ToEqual("nice_uri")
  spec.Expect(bv.Data).ToEqual("nice_data")
}

func TestSaveInsertNewBeatVector(t *testing.T) {
  spec := gspec.New(t)
  tests.DbCleanUp("beatvectors")
  e := Builder().Uri("awesome_uri").Data("nice_data").Bv.Save()
  if e != nil { t.Fatal("failed to save: ", e) }
  bv := Find("awesome_uri")
  spec.Expect(bv.Uri).ToEqual("awesome_uri")
  spec.Expect(bv.Data).ToEqual("nice_data")
}

func TestUpdatesExistingBeatVector(t *testing.T) {
  tests.DbCleanUp("beatvectors")
  insertFakeBeatVector()
  bv := Find("nice_uri")
  bv.Data = "not_so_awesome_data"
  bv.Save()
  gspec.New(t).Expect(Find("nice_uri").Data).ToEqual("not_so_awesome_data")
}

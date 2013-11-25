package beatvectors

import (
  "testing"
  "beatstorm/core"
  "beatstorm/tests"
  "github.com/cviedmai/gspec"
  "beatstorm/models/beatvector"
  "github.com/garyburd/redigo/redis"
)

var (
  originalGetBv = getBv
  originalMarkBv = markBv
)

func TestWorksIfFound(t *testing.T) {
  bv := beatvector.Builder().Uri("an_uri").Data("some_data").Bv
  stubGetBv(bv)
  actual := Show(tests.CB().ParamId("an_uri").C)
  tests.AssertBody(t, `{"uri":"an_uri","data":"some_data"}`, actual)
  tests.AssertStatus(t, 200, actual)
}

func TestStoresUriToProcessIfNotFound(t *testing.T) {
  tests.DbCleanUp("beatvectors")
  finished := make(chan bool)
  spyMarkBv(finished)
  actual := Show(tests.CB().ParamId("an_uri").C)
  tests.AssertStatus(t, 404, actual)
  <- finished
  queued, _ := redis.Bool(tests.Redis().Do("sismember", "q:bv", "an_uri"))
  gspec.New(t).Expect(queued).ToEqual(true)
}

func stubGetBv(bv core.BeatVector) {
  getBv = func(uri string) core.BeatVector {
    getBv = originalGetBv
    return bv
  }
}

func spyMarkBv(c chan bool) {
  markBv = func (uri string) {
    originalMarkBv(uri)
    c <- true
    markBv = originalMarkBv
  }
}

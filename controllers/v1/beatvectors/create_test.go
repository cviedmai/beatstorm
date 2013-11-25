package beatvectors

import (
  "testing"
  "beatstorm/core"
  "beatstorm/tests"
  "github.com/cviedmai/auwfg"
  "github.com/cviedmai/gspec"
  "beatstorm/models/beatvector"
)

var (
  originalAuthenticate = authenticate
)

func TestErrorOnCreateWhenNoPermission(t *testing.T) {
  actual := Create(tests.CB().C)
  gspec.New(t).Expect(actual.GetStatus()).ToEqual(401)
}

func TestErrorOnCreateWhenDataIsMissing(t *testing.T) {
  spec := gspec.New(t)
  stubAuthenticate()
  actual := Create(tests.CB().Body(&CreateInput{}).C)
  spec.Expect(string(actual.GetBody())).ToEqual(`{"uri":["uri is required"],"data":["data is required"]}`)
  spec.Expect(actual.GetStatus()).ToEqual(400)
}

func TestWorksWhenAllIsFine(t *testing.T) {
  tests.DbCleanUp("beatvectors")
  stubAuthenticate()
  actual := Create(tests.CB().Body(&CreateInput{Uri: "abc", Data: "ced"}).C)
  bv := beatvector.Find("abc")
  if bv == nil { t.Fatal("Beatvector was not created") }
  gspec.New(t).Expect(bv.Data).ToEqual("ced")
  tests.AssertResponse(t, core.CreatedResponse, actual)
}

func stubAuthenticate() {
  authenticate = func(c *core.Context) auwfg.Response {
    authenticate = originalAuthenticate
    return nil
  }
}

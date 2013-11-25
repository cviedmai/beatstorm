package tests

import (
  "testing"
  "github.com/cviedmai/gspec"
  "github.com/cviedmai/auwfg"
)

var DummyResponse = auwfg.Json("whatever").Status(6555)

func AssertResponse(t *testing.T, expected auwfg.Response, actual auwfg.Response) {
  spec := gspec.New(t)
  spec.Expect(actual.GetStatus()).ToEqual(expected.GetStatus())
  spec.Expect(string(actual.GetBody())).ToEqual(string(expected.GetBody()))
}

func AssertBody(t *testing.T, expected string, actual auwfg.Response) {
  spec := gspec.New(t)
  spec.Expect(string(actual.GetBody())).ToEqual(expected)
}

func AssertStatus(t *testing.T, expected int, actual auwfg.Response) {
  spec := gspec.New(t)
  spec.Expect(actual.GetStatus()).ToEqual(expected)
}

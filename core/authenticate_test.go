package core

import (
  "testing"
  "github.com/cviedmai/auwfg"
  "github.com/cviedmai/gspec"
)

func TestValidSig(t *testing.T) {
  context := NewContext(&auwfg.BaseContext{
    RawPath: "/beatvector",
    RawQuery: "a=1&b=2",
    RawBody: []byte(`{"data":"a lot of it"}`),
  })

  vs := validSig(context, "9e99486e5e1917d6a8d707d1b3bb7f48ff88575a")
  gspec.New(t).Expect(vs).ToEqual(true)
}

package tests

import (
  "beatstorm/core"
  "github.com/cviedmai/auwfg"
)

type ContextBuilder struct {
  C *core.Context
}

func CB() *ContextBuilder {
  return &ContextBuilder {
    C: &core.Context{
      BaseContext: &auwfg.BaseContext{
        Params: new(auwfg.Params),
      },
    },
  }
}

func (b *ContextBuilder) Body(body interface{}) *ContextBuilder {
  b.C.Body = body
  return b
}

func (b *ContextBuilder) ParamId(id string) *ContextBuilder {
  b.C.Params.Id = id
  return b
}

func (b *ContextBuilder) Query(key, value string) *ContextBuilder {
  if b.C.Query == nil { b.C.Query = make(map[string]string) }
  b.C.Query[key] = value
  return b
}

func (b *ContextBuilder) RawBody(rb []byte) *ContextBuilder {
  b.C.RawBody = rb
  return b
}

func (b *ContextBuilder) RawQuery(q string) *ContextBuilder {
  b.C.RawQuery = q
  return b
}

func (b *ContextBuilder) RawPath(p string) *ContextBuilder {
  b.C.RawPath = p
  return b
}

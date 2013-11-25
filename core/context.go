package core

import (
  "github.com/cviedmai/auwfg"
)

type Context struct {
  *auwfg.BaseContext
}

func NewContext(bc *auwfg.BaseContext) *Context {
  return &Context{bc}
}

package server

import (
  "github.com/cviedmai/auwfg"
  "beatstorm/controllers/v1/beatvectors"
)

func addRoutes(c *auwfg.Configuration) {
  c.Route(auwfg.R("GET", "v1", "beatvectors", beatvectors.Show))
  c.Route(auwfg.R("POST", "v1", "beatvectors", beatvectors.Create).BodyFactory(func() interface{} { return new(beatvectors.CreateInput) } ))
}

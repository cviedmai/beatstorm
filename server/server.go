package server

import (
  "fmt"
  "time"
  "strings"
  "runtime"
  "beatstorm/core"
  "github.com/cviedmai/auwfg"
)

func Run() {
  runtime.GOMAXPROCS(runtime.NumCPU())
  config := core.GetConfig()
  setupSignals()
  startStats(config.StatsFile)

  auwfgConfig := auwfg.Configure().Address(config.Address).ContextFactory(contextFactory).Dispatcher(dispatcher).LoadRawBody().LoadRawQuery().LoadRawPath()
  addRoutes(auwfgConfig)
  fmt.Println(time.Now().String() + " - BeatStorm running " + strings.ToUpper(core.GetEnv()) + " mode on " + config.Address)
  auwfg.Run(auwfgConfig)
}

func dispatcher(route *auwfg.Route, context interface{}) auwfg.Response {
  return route.Action.(func(*core.Context) auwfg.Response)(context.(*core.Context))
}

func contextFactory(bc *auwfg.BaseContext) interface{} {
  return core.NewContext(bc)
}

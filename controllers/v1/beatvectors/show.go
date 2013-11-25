package beatvectors

import(
  "log"
  "beatstorm/core"
  "beatstorm/core/redis"
  "github.com/cviedmai/auwfg"
  "beatstorm/models/beatvector"
)

var (
  getBv = func(uri string) core.BeatVector {
    bv := beatvector.Find(uri)
    if bv == nil { return nil }
    return bv
  }
)
func Show(context *core.Context) auwfg.Response {
  bv := getBv(context.Params.Id)
  if bv != nil { return auwfg.Json(bv.Json()) }
  go markBv(context.Params.Id)
  return core.NotFoundResponse
}

var markBv = func(uri string) {
  r := redis.Db.Get()
  defer r.Close()
  if _, err := r.Do("sadd", "q:bv", uri); err != nil {
    log.Println("Saving in redis: ", err)
  }
}

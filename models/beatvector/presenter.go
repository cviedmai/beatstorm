package beatvector

import (
  "time"
  "beatstorm/core/cache"
)

func (bv *BeatVector) Json() string {
  json, _ := cache.FetchString("bv:" + bv.Uri + ":json", time.Minute * 30, func() (string, error) {
    return bv.json(), nil
  })
  return json
}

func (bv *BeatVector) json() string {
  return `{"uri":"` + bv.Uri + `","data":"` + bv.Data + `"}`
}

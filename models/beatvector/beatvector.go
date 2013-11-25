package beatvector

import (
  "log"
  "time"
  "errors"
  "beatstorm/core/cache"
  "beatstorm/core/postgres"
)

type BeatVector struct {
  Uri string
  Data string
}

var (
  CantSaveError = errors.New("Cannot save the user data")
)

const (
  CacheDuration = time.Minute * 10
  SelectQuery = "select data from beatvectors where uri = $1"
  UpdateQuery = "update beatvectors set data=$2 where uri=$1"
  InsertQuery = "insert into beatvectors (uri, data) values ($1, $2)"
)

func New() *BeatVector {
  return &BeatVector{}
}

func Find(uri string) *BeatVector {
  item, err := cache.Fetch(cacheKey(uri), CacheDuration, func() (interface{}, error) {
    rows, err := postgres.Db.Query(SelectQuery, uri)
    if err != nil { return nil, err }
    defer rows.Close()
    if found := rows.Next(); !found { return nil, nil }
    bv := New()
    bv.Uri = uri
    rows.Scan(&bv.Data)
    return bv, nil
  })
  if item != nil { return item.(*BeatVector) }
  if err != nil { log.Println("Retrieving data: ", err) }
  return nil
}

func (bv *BeatVector) Save() error {
  res, err := postgres.Db.Exec(UpdateQuery, bv.Uri, bv.Data)
  if err != nil { return err }
  if n, err := res.RowsAffected(); err != nil {
    return err
  } else if n == 0 {
    // nothing was updated, lets try saving it
    res, err := postgres.Db.Exec(InsertQuery, bv.Uri, bv.Data)
    if err != nil { return err }
    if n, err := res.RowsAffected(); err != nil {
      return err
    } else if n == 0 {
      return CantSaveError
    }
  }

  cache.Set(cacheKey(bv.Uri), bv, CacheDuration)
  return nil
}

func cacheKey(uri string) string {
  return "bv:" + uri
}

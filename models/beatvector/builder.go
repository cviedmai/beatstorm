package beatvector

import (
  "beatstorm/core/postgres"
)

type B struct {
  Bv *BeatVector
}

func Builder() *B {
  Bv := New()
  return &B{Bv}
}

func (b *B) Uri(uri string) *B { b.Bv.Uri = uri; return b }
func (b *B) Data(data string) *B { b.Bv.Data = data; return b }

func insertFakeBeatVector() {
  postgres.Db.Exec(`insert into beatvectors (uri, data) values ('nice_uri', 'nice_data');`)
}

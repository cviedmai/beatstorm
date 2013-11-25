package postgres

import (
  "testing"
  "github.com/cviedmai/gspec"
)

func TestGetItems(t *testing.T) {
  spec := gspec.New(t)
  rows, _ := Db.Query("select 'a'")
  rows.Next()
  var x string
  rows.Scan(&x)
  spec.Expect(x).ToEqual("a")
}

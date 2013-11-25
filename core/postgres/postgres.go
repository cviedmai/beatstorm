package postgres

import (
  "log"
  "database/sql"
  "beatstorm/core"
  _ "github.com/lib/pq"
)

type DB struct {
  conn *sql.DB
}

var (
  Db *DB
  throttle chan bool
)

func init() {
  config := core.GetConfig().Postgres
  conn, err := sql.Open("postgres", config.Config)
  if err != nil { log.Fatal(err) }
  conn.SetMaxIdleConns(1)
  throttle = make(chan bool, config.PoolSize)
  Db = &DB{conn: conn,}
}

func (d *DB) Query(sql string, params ...interface{}) (*sql.Rows, error) {
  throttle <- true
  defer func() { <- throttle }()
  if len(params) == 0 { return d.conn.Query(sql) }
  return d.conn.Query(sql, params...)
}

func (d *DB) Exec(sql string, params ...interface{}) (sql.Result, error) {
  throttle <- true
  defer func() { <- throttle }()
  if len(params) == 0 { return d.conn.Exec(sql) }
  return d.conn.Exec(sql, params...)
}

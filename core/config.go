package core

import(
  "os"
  "log"
  "strings"
  "io/ioutil"
  "encoding/json"
)

const (
  TEST = "test"
  PRODUCTION = "production"
  DEVELOPMENT = "development"
)

type RedisConfig struct {
  Db int
  Host string
  PoolSize int
}

type PostgresConfig struct {
  Config string
  PoolSize int
}

type Config struct {
  Env string
  Secret string
  Address string
  StatsFile string
  Redis RedisConfig
  Postgres PostgresConfig
}

var (
  config *Config
)

func GetConfig() *Config {
  return config
}

func init() {
  loadConfig()
}

func loadConfig() {
  configPath := GetPath("config.json")
  if _, err := os.Stat(configPath); err != nil {
    log.Fatal("Cannot access config file")
  }

  file, err := ioutil.ReadFile(configPath)
  if err != nil { log.Fatal(err) }

  temp := make(map[string]*Config)
  if err = json.Unmarshal(file, &temp); err != nil {
    log.Fatal("Unmarshaling config json: ", err)
  }

  env := GetEnv()
  config = temp[env]
  config.Env = env
}

func GetEnv() string {
  env := os.Getenv("BS_ENV")
  for _, validEnv := range []string{TEST, DEVELOPMENT, PRODUCTION} {
    if env == validEnv { return env }
  }

  if len(os.Args) >= 1 && strings.Contains(os.Args[0], ".test") { return TEST }
  return DEVELOPMENT
}

package core

type BeatVector interface {
  Save() error
  Json() string
}

package main

import (
  "time"
)

type Sequence struct {
  Name string `json:"name"`
  User string `json:"user"`
  Creation time.Time `json:"creation"`
  Scenes []Scene `json:"scenes"`
}

type Location []Sequence

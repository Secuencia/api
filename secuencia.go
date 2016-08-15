package main

import (
  "time"
)

type Sequence struct {
  Id int `json:"id"`
  Name string `json:"name"`
  User string `json:"user"`
  Country Place `json:"country"`
  City Place `json:"city"`
  Creation time.Time `json:"creation"`
  Scenes []Scene `json:"scenes"`
}

type Place string

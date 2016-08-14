 package main

 import ()

 type Scene struct {
   Name string `json:"name"`
   Description string `json:"description"`
   Coordinates Location
   Tags []string `json:"tags"`
 }

type Location struct {
  Latitude string `json:"latitude"`
  Longitude string `json:"longitude"`
}

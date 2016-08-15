// In charge to handle the end points.

package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "html"

  "github.com/gorilla/mux"
)


func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func SequencesIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json;charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(sequences); err != nil {
    panic(err)
  }
}

func GetSequence(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fmt.Fprintln(w, "This is the sequence id: ",vars["sequenceId"])
}

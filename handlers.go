// In charge to handle the end points.

package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "html"
  "io"
  "io/ioutil"

  "github.com/gorilla/mux"
)

// REVIEW: Description
func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// REVIEW: Description
func SequencesIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","application/json;charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(sequences); err != nil {
    panic(err)
  }
}

// REVIEW: Description
func GetSequence(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fmt.Fprintln(w, "This is the sequence id: ",vars["sequenceId"])
}

// REVIEW: Description
func CreateSequence(w http.ResponseWriter, r *http.Request) {
  var sequence Sequence
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

  if err != nil {
    panic(err)
  }

  if err := r.Body.Close() ; err != nil {
    panic(err)
  }

  if err := json.Unmarshal(body, &sequence); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422) // This header corresponds to "unprocessable entity"
    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  s := RepoCreateSequence(sequence)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)
  if err := json.NewEncoder(w).Encode(s); err != nil {
    panic(err)
  }
}

// Try the above function with the following json in Postman. Dont forget to configure the Content-Type headers.

//{
// "name":"Test of a new sequence",
// "user":"chadepl",
// "country":"Colombia",
// "city":"Bogota",
// "scenes":[
//     {"name":"Scene 1","description":"Description 1."},
//     {"name":"Scene 2","description":"Description 2. "}
//     ]
//}

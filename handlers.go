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
  sequences := []Sequence{
    Sequence{
      Name: "Around Bogota",
      User: "mapus",
      Scenes: []Scene{
        Scene{
          Name: "Monserrate",
          Description: "One of the most traditional places to go.",
        },
        Scene{
          Name: "La Puerta Falsa",
          Description: "Try the Tamal, a traditional colombian dish.",
        },
      },
    },
    Sequence{
      Name: "Top 10 burgers Bogota",
      User: "Tulio",
      Scenes: []Scene{
        Scene{
          Name: "El Gordo ",
          Description: "Simulates a Brooklyn bistro. Amazing flavor.",
        },
        Scene{
          Name: "El Corral",
          Description: "To Colombia, from Colombia. A colombian classic.",
        },
        Scene{
          Name: "Home Burgers",
          Description: "Try the classic US burger here in Colombia.",
        },
      },
    },
  }

  json.NewEncoder(w).Encode(sequences)
}

func GetSequence(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fmt.Fprintln(w, "This is the sequence id: ",vars["sequenceId"])
}

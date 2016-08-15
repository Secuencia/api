// Mock database until a fully functional one is added

package main

import "fmt"


var currentId int

var sequences []Sequence

// Get some initial data
func init() {
  RepoCreateSequence(
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
  )
  RepoCreateSequence(
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
  )
}

func RepoFindSequence(sequenceId int) Sequence {
  for _, s := range sequences {
    if s.Id == sequenceId {
      return s
    }
  }
  // Return empty sequence if not found
  return Sequence{}
}

func RepoCreateSequence(s Sequence) Sequence {
  currentId += 1
  s.Id = currentId
  sequences = append(sequences, s)
  return s
}

func RepoDestroySequence(sequenceId int) error {
  for i, s := range sequences {
    if s.Id == sequenceId {
      sequences = append(sequences[:i], sequences[i+1:]...)
      return nil
    }
  }
  return fmt.Errorf("Could not find the sequence with id %d to delete", sequenceId)
}

package main

import (
  "log"
  "os"

  "github.com/dotabuff/manta"
  "github.com/dotabuff/manta/dota"
)

func main() {
  // Create a new parser instance from a file. Alternatively see NewParser([]byte)
  f, err := os.Open("replays//6299976009.dem")
  if err != nil {
    log.Fatalf("unable to open file: %s", err)
  }
  defer f.Close()

  p, err := manta.NewStreamParser(f)
  if err != nil {
    log.Fatalf("unable to create parser: %s", err)
  }

  // Register callbacks
  p.Callbacks.OnCDOTAUserMsg_ChatEvent(func (m *dota.CDOTAUserMsg_ChatEvent) error {
    log.Printf("%s\n", m)
    return nil
  })

  // Start parsing the replay!
  p.Start()

  log.Printf("Parse Complete!\n")
}

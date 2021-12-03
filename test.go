package main

import(
  "log"
  "os"

  "github.com/dotabuff/manta"
  "github.com/dotabuff/manta/dota"
)

func main() {
  // Open replay file
  f, err := os.Open("./replays/6299946143.dem")
  if err != nil {
    log.Fatalf("unable to open file: %s", err)
  }
  defer f.Close()

  // Create parser from replay file
  p, err := manta.NewStreamParser(f)
  if err != nil {
    log.Fatalf("unable to create parser: %s", err)
  }

  entity_count := uint32(0)
  p.OnEntity(func(e *manta.Entity, o manta.EntityOp) error {
    entity_count += 1
    log.Printf("%v | e. type: %T", e.GetClassName(), e)
    return nil
  })

  p.Callbacks.OnCDOTAUserMsg_ChatMessage(func (m* dota.CDOTAUserMsg_ChatMessage) error {
    log.Printf("%s", m)
    return nil
  })

  p.Start()
  log.Printf("Parse complete!")
  //log.Printf("Got %v entity events!", entity_count)
}

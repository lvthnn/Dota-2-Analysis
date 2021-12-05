package main

import(
  "log"
  "os"
  "strings"
  "strconv"
  "time"

  "github.com/dotabuff/manta"
  //"github.com/dotabuff/manta/dota"
)

var preGameStartTime time.Duration
var gameTime time.Duration
var gameStartTime time.Duration
var gameEndTime time.Duration

var heroID [10]int32 

func main() {
  // Open replay file
  f, err := os.Open("./replays/6300048284.dem")
  if err != nil {
    log.Fatalf("unable to open file: %s", err)
  }
  defer f.Close()

  // Create parser from replay file
  p, err := manta.NewStreamParser(f)
  if err != nil {
    log.Fatalf("unable to create parser: %s", err)
  }

  p.OnEntity(func(e *manta.Entity, o manta.EntityOp) error {
   
    if e.GetClassName() == "CDOTAGamerulesProxy" {
      pt, _ := e.GetFloat32("m_pGameRules.m_flPreGameStartTime")
      preGameStartTime = time.Duration(pt) * time.Second

      t, _ := e.GetFloat32("m_pGameRules.m_fGameTime")
      gameTime = time.Duration(t) * time.Second
      
      ts, _ := e.GetFloat32("m_pGameRules.m_flGameStartTime")
      gameStartTime = time.Duration(ts) * time.Second

      te, _ := e.GetFloat32("m_pGameRules.m_flGameEndTime")
      gameEndTime = time.Duration(te) * time.Second
    }

    if e.GetClassName() == "CDOTA_PlayerResource" {
      for i := 0; i < 10; i++ {
        heroID[i], _ = e.GetInt32("m_vecPlayerTeamData.000" + strconv.Itoa(i) + ".m_nSelectedHeroID")
      } 
    }

    if strings.HasPrefix(e.GetClassName(), "CDOTA_Unit_Hero") {
      x, _ := e.GetFloat32("CBodyComponent.m_vecX")
      y, _ := e.GetFloat32("CBodyComponent.m_vecY")
      owner, _ := e.GetInt32("m_iPlayerID")
      log.Printf("time: %v hero: %v x: %v y: %v owner: %v", gameTime, heroID[owner], x, y, owner)
    }

    return nil
  })

  p.Start()
  log.Printf("Parse complete!")
}

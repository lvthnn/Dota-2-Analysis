package main

import(
  "log"
  "os"
  "strings"
  "strconv"
  "encoding/csv"
  "fmt"

  "github.com/dotabuff/manta"
  //"github.com/dotabuff/manta/dota"
)

var preGameStartTime float32 
var gameTime float32
var gameStartTime float32
var gameEndTime float32 

var gameWinner int32
var gameMode int32

var heroID [10]int32 
var netWorth [10]int32
var totalXP [10]int32

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

  w, err := os.Create("./output.csv")
  if err != nil {
    log.Fatal("cannot create output file: %s", err)
  }
  defer w.Close()

  writer := csv.NewWriter(w)
  defer writer.Flush()

  outputHeaders := []string{
    "preGameStartTime", 
    "gameTime", 
    "heroID", 
    "cellX", 
    "cellY", 
    "vecX",
    "vecY",
    "playerID",
    "team",
    "netWorth",
    "totalXP",
    "gameWinner",
    "gameMode",
  }
  writer.Write(outputHeaders)

  p.OnEntity(func(e *manta.Entity, o manta.EntityOp) error {
   
    if e.GetClassName() == "CDOTAGamerulesProxy" {
      gameWinner, _ = e.GetInt32("m_pGameRules.m_nGameWinner")
      gameMode, _ = e.GetInt32("m_pGameRules.m_iGameMode")

      preGameStartTime, _ = e.GetFloat32("m_pGameRules.m_flPreGameStartTime")
      gameTime, _ = e.GetFloat32("m_pGameRules.m_fGameTime")
      gameStartTime, _ = e.GetFloat32("m_pGameRules.m_flGameStartTime")
      gameEndTime, _ = e.GetFloat32("m_pGameRules.m_flGameEndTime")
    }

    if e.GetClassName() == "CDOTA_PlayerResource" {
      for i := 0; i < 10; i++ {
        heroID[i], _ = e.GetInt32("m_vecPlayerTeamData.000" + strconv.Itoa(i) + ".m_nSelectedHeroID")
      } 
    }
  
    if e.GetClassName() == "CDOTA_DataDire" || e.GetClassName() == "CDOTA_DataRadiant" {
      for i := 0; i < 5; i++ {
        numPlayer := i
        if(e.GetClassName() == "CDOTA_DataDire") { numPlayer += 5 }

        nw, _ := e.GetInt32("m_vecDataTeam.000" + strconv.Itoa(i) + ".m_iNetWorth")
        xp, _ := e.GetInt32("m_vecDataTeam.000" + strconv.Itoa(i) + ".m_iTotalEarnedXP")
        netWorth[numPlayer] = nw
        totalXP[numPlayer] = xp
      }
    }


    if strings.HasPrefix(e.GetClassName(), "CDOTA_Unit_Hero") {
      cellX, _ := e.GetUint64("CBodyComponent.m_cellX")
      cellY, _ := e.GetUint64("CBodyComponent.m_cellY")
      vecX, _ := e.GetFloat32("CBodyComponent.m_vecX")
      vecY, _ := e.GetFloat32("CBodyComponent.m_vecY")
      owner, _ := e.GetInt32("m_iPlayerID")
      team, _ := e.GetUint64("m_iTeamNum")

      output := []string{
        fmt.Sprintf("%.4f", preGameStartTime), 
        fmt.Sprintf("%.4f", gameTime),
        strconv.Itoa(int(heroID[owner])), 
        strconv.Itoa(int(cellX)),
        strconv.Itoa(int(cellY)),
        fmt.Sprintf("%.4f", vecX), 
        fmt.Sprintf("%.4f", vecY), 
        strconv.Itoa(int(owner)),
        strconv.Itoa(int(team)),
        strconv.Itoa(int(netWorth[owner])),
        strconv.Itoa(int(totalXP[owner])),
        strconv.Itoa(int(gameWinner)),
        strconv.Itoa(int(gameMode)),
      }
      writer.Write(output)
    }
    return nil
  })

  p.Start()
  log.Printf("Parse complete!")
  f.Close()
}

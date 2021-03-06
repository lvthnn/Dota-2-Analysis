package main 

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

  "github.com/dotabuff/manta"
  "github.com/ermanimer/progress_bar"
)

// OTHER VARIABLES
var filePaths []string
var fileNames []string

// PARSER VARIABLES

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
  fmt.Printf("Starting replay parse...")

  // Loop through replays directory and append all matching files
  // to filePaths and fileNames arrays
	err := filepath.Walk("./replays/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("%s", err)
			return err
		}

		if !info.IsDir() {
			filePaths = append(filePaths, path)
			fileNames = append(fileNames, info.Name())
		}
		return nil
	})

	if err != nil {
		log.Fatalf("%s", err)
	}

  // Instantiate progress bar
  output := os.Stdout
  schema := "({bar}) ({current} of {total} completed)"
	filledCharacter := "▆"
	blankCharacter := " "
  
  numReplays := len(fileNames)
	var length float64 = float64(numReplays)
	var totalValue float64 = float64(numReplays) 

  pb := progress_bar.NewProgressBar(output, schema, filledCharacter, blankCharacter, length, totalValue)
  pb.Start()

  // Parse all replays found during scan. Update progress bar for each
  // replay parse completed by the script
	for i := 0; i < len(fileNames); i++ {
		parse(filePaths[i], fileNames[i])
    pb.Update(float64(i + 1))
	}
}

func parse(path string, name string) {
	// Open replay file
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("unable to open file: %s", err)
	}
	defer f.Close()

	// Create parser from replay file
	p, err := manta.NewStreamParser(f)
	if err != nil {
		log.Fatalf("unable to create parser: %s", err)
	}

	w, err := os.Create("./data/output_" + strings.ReplaceAll(name, ".dem", "") + ".csv")
	if err != nil {
		log.Fatalf("cannot create output file: %s", err)
	}
	defer w.Close()

	writer := csv.NewWriter(w)
	defer writer.Flush()

	outputHeaders := []string{
		"preGameStartTime",
		"gameTime",
		"gameEndTime",
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
				if e.GetClassName() == "CDOTA_DataDire" {
					numPlayer += 5
				}

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
				fmt.Sprintf("%.4f", gameEndTime),
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
	f.Close()

}

package main

import(
  "log"
  "os"
  "io"
  "net/http"
  "errors"
)

func main() {
  // reqFiles array stores fileName and URL
  reqFiles := [3][2]string{
    {"tables/npc_heroes.txt", "https://github.com/SteamDatabase/GameTracking-Dota2/raw/master/game/dota/scripts/npc/npc_heroes.txt"},
    {"tables/npc_abilities.txt", "https://raw.githubusercontent.com/SteamDatabase/GameTracking-Dota2/master/game/dota/scripts/npc/npc_abilities.txt"},
    {"tables/npc_units.txt", "https://raw.githubusercontent.com/SteamDatabase/GameTracking-Dota2/master/game/dota/scripts/npc/npc_units.txt"},
  }

  for i := 0; i < len(reqFiles); i++ {
    pullTable(reqFiles[i][0], reqFiles[i][1])
  } 

}

func pullTable(fileName, URL string) error {
  response, err := http.Get(URL)
  if err != nil {
    return err
  }

  log.Printf("Request URL: %s successful", URL)
  defer response.Body.Close()

  if response.StatusCode != 200 {
    return errors.New("Recieved non-200 response code")
  }

  file, err := os.Create(fileName)
  if err != nil {
    return err
  }

  log.Printf("Created file %s", fileName)

  _, err = io.Copy(file, response.Body)
  if err != nil {
    return err
  }

  log.Printf("Finished writing to file!")

  return nil

}

func parseRegex(fileName string) error {
  return nil
}

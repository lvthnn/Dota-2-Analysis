# Dota 2 Data Parser
Uses Dotabuff's [Manta parser](https://github.com/dotabuff/manta) to extract match data from game replays.
Mainly useful for spatiotemporal analysis of data. Recommend using [OpenDota API](https://docs.opendota.com/#) for statistical analysis of
non-spatiotemporal factors (e.g. hero drafts).
To start, clone the Github repository into the directory of your choice:

```git clone https://github.com/lvthnn/Dota2_DataMining```

Place the replays you want to parse into the `replays` folder and run the `parse.go` script. The script returns a parsed
`.csv` file which can be found in the `data` folder. The `tables` folder contains `npc_heroes.txt`, `npc_abilities.txt`
and `npc_units.txt` pulled from the [Dota 2 Gametracking](https://github.com/SteamDatabase/GameTracking-Dota2) repository.
To fetch the tables, run the command:

```go run gettables.go```

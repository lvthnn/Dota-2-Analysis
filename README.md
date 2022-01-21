# Dota 2 Data Parser
Uses Dotabuff's [Manta parser](https://github.com/dotabuff/manta) to extract match data from game replays.
Mainly useful for spatiotemporal analysis of data. Recommend using OpenDota API for statistical analysis of
non-spatiotemporal factors (e.g. hero drafts).
To start, clone the Github repository into the directory of your choice:

```git clone https://github.com/lvthnn/Dota2_DataMining```

Place the replays you want to parse into the `replays` folder and run the script. The script returns a parsed
`.csv` file which can be found in the `data` folder.

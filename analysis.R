library("rjson")
heroes_data <- fromJSON(file = "~/Github/Dota2_DataMining/tables/npc_heroes.json")

setwd("~/Github/Dota2_DataMining/data/")
files <- list.files(path = "~/Github/Dota2_DataMining/data/")
dataframes <- lapply(files, read.csv)
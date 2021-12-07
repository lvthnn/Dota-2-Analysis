# Imports
library(ggplot2, dplyr, plotly)

# Read in CSV output file from repo
f <- file.path("~/Github/Dota2_DataMining/output.csv")
df <- read.csv(f)

# Notice: Team 2 = Radiant
#         Team 3 = Dire

# Victory codes: 2 = Radiant
#                3 = Dire
#                5 = Undecided
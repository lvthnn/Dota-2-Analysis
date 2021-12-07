# Imports
library(ggplot2, dplyr, plotly)

# Read in CSV output file from repo
f <- file.path("~/Github/Dota2_DataMining/output.csv")
df <- read.csv(f)

# To see the total movement path of
# heroes, we use ggplot visualization:
df %>%
  ggplot(aes(x = cellX, y = cellY)) + 
  geom_path() +
  theme_minimal()
  facet_wrap(~ heroID)

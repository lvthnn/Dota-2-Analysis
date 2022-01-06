# Imports
library(ggplot2, dplyr)

# Read in CSV output file from repo
f <- file.path("~/Github/Dota2_DataMining/output.csv")
df <- read.csv(f)

# _____________________________
#|         |                   |
#| Notice  | Team 2 = Radiant  |
#|         | Team 3 = Dire     |
#|_________|___________________|
#|         |                   |
#| Victory | 2 = Radiant       |
#|   codes | 3 = Dire          |
#|         | 5 = Undecided     |
#|_________|___________________|

# Plot libraries
ggplot(df, aes(x = cellX, y = cellY, color = gameTime)) +
  geom_path() +
  facet_wrap(~ playerID) +
  theme_minimal() +
  theme(legend.position = "bottom")

df %>%
  filter(gameTime == 802.5590) %>%
  ggplot(aes(x = cellX, y = cellY, color = as.factor(playerID))) +
  geom_point()
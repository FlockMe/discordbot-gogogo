import (
  "github.com/bwmarrin/discordgo"
  "os"
  "log"
)

token := os.Getenv("BOT_TOKEN")

if token == "" {
 return
}

// Bot session

dg, err := discordgo.New(token)

if err != nil {
 log.Println(err)
 return
}

// New Application

app := &discordgo.Application{}
app.Name = "Nico Nico"
app.Description = "Nico nico is a bot."

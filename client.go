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

app, err := dg.ApplicationCreate(app)
log.Printf("ApplicationCreate: err: %+v, app: %+v\n", err, ap)

// Create the bot account

bot, err := dg.ApplicationBotCreate(app.ID)
log.Printf("BotCreate: err: %+v, bot: %+v\n", err, bot)


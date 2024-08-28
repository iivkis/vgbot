package main

import (
	"github.com/iivkis/vgbot"
)

func main() {
	botapi := vgbot.NewVKBotAPI("")

	r := vgbot.NewRouter()

	r.IncludeAdjuster(
		NewBasicHandler(botapi),
	)
}

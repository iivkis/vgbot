package main

import (
	"github.com/iivkis/vgbot"
)

func main() {
	vkapi := vgbot.NewVKAPI(vgbot.NewVKAPIProvider(""))

	r := vgbot.NewRouter()

	r.IncludeHandler(
		NewBasicHandler(vkapi),
	)
}

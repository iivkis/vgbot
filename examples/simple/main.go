package main

import (
	"github.com/iivkis/vgbot"
)

func main() {
	vkapi := vgbot.NewVKAPI("")

	r := vgbot.NewRouter()

	r.IncludeHandler(
		NewBasicHandler(vkapi),
	)
}

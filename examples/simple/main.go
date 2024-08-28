package main

import (
	"github.com/iivkis/vgbot"
)

func main() {
	r := vgbot.NewRouter()

	r.IncludeAdjuster(
		NewBasicHandler(),
	)
}

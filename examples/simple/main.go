package main

import (
	"fmt"

	"github.com/iivkis/vgbot"
)

func main() {
	fmt.Println("123")

	vkapi := vgbot.NewVKAPI(vgbot.NewVKAPIProvider("vk1.a.u8geTFNzBinJ6_AgOGqfvzDBUqCFdZswQJ2lC47hV9lNbEo-DYz8-UvEbCWaKF8WwRid6Y8wgx_DYjCq4qatNs9zwtEPYcNso_hU5prNcsiXf4fD-LCvlMXk6MyYkh89_rVVcjY2NY0CjwoT6ytv13EKNpbg5B7kKF12UE9rvqsDm2gCCrZXdZ7zPDZ0K75R1_r62R9_QW5ZFlatIbSBcw"))
	fmt.Println(vkapi.Groups().GetLongPollServer(182948792))

	r := vgbot.NewRouter()

	r.IncludeHandler(
		NewBasicHandler(vkapi),
	)
}

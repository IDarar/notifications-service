package main

import "github.com/IDarar/notifications-service/internal/app"

const configPath = "configs/main"

func main() {
	app.Run(configPath)
}

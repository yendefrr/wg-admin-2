package main

import "github.com/yendefrr/wg-admin/internal/app"

const configPath = "configs/main"

func main() {
	app.Run(configPath)
}

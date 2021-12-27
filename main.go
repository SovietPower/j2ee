package main

import (
	"j2ee/config"
	"j2ee/router"
)

func main() {
	config.Init()
	r := router.NewRouter()
	r.Run(":8080")
}

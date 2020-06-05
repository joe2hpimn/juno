package main

import (
	"fmt"
	"log"

	"github.com/douyu/juno/router"
	"github.com/douyu/jupiter/pkg/conf"
)

func main() {
	app := router.New()
	app.SetGovernor(fmt.Sprintf("0.0.0.0:%d", conf.GetInt("jupiter.server.govern.port")))
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"myproject/pkg/router"
)

func main() {
    r := router.SetupRouter()
    r.Run(":8080")
}

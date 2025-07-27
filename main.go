package main

import (
	"github.com/RicardoSantos-99/payment_dispatcher_go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":9999")
}

package routes

import (
	"github.com/RicardoSantos-99/payment_dispatcher_go/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/payments-summary", handlers.SummaryPayments)
	r.POST("/payment", handlers.CreatePayment)
	r.POST("/purge-payments", handlers.PurgePayments)
}

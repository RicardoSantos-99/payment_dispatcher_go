package handlers

import (
	"net/http"

	"github.com/RicardoSantos-99/payment_dispatcher_go/models"
	"github.com/gin-gonic/gin"
)

func SummaryPayments(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"default":  gin.H{"totalAmount": 0, "totalRequests": 0},
		"fallback": gin.H{"totalAmount": 0, "totalRequests": 0},
	})
}

func CreatePayment(context *gin.Context) {
	var payment models.Payment

	if err := context.ShouldBindJSON(&payment); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "payment received", "amount": payment.Amount})
}

func PurgePayments(context *gin.Context) {

}

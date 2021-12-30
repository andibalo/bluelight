package handler

import (
	"bluelight/helper"
	"bluelight/transaction"
	"bluelight/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.GetTransactionsInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Failed to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)

	if err != nil {
		response := helper.APIResponse("Failed to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("campaigns transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(userID)

	if err != nil {
		response := helper.APIResponse("Failed to get user transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("user transactions", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {

	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create transaction failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newTransaction, err := h.service.CreateTransaction(input)

	if err != nil {
		response := helper.APIResponse("Create transaction failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("New user transaction", http.StatusOK, "success", transaction.FormatTransaction(newTransaction))

	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input transaction.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.APIResponse("Failed to get notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.ProcessPayment(input)

	if err != nil {
		response := helper.APIResponse("Failed to get notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("New user transaction", http.StatusOK, "success", nil)

	c.JSON(http.StatusOK, response)
}

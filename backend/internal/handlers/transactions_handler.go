package handlers

import (
	"backend/internal/dto"
	"backend/internal/interfaces"
	"backend/internal/mappers"
	apperrors "backend/pkg/appErrors"
	"encoding/json"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

var (
	ErrTransactionHandlerUnprocessableEntity = apperrors.ErrUnprocessableEntity("Transaction")
	ErrTransactionHandlerNotFound            = apperrors.ErrNotFound("Transaction")
	ErrTransactionInternalServerError        = apperrors.ErrInternalServerError("Transaction")
)

type TransactionHandler struct {
	TransactionService interfaces.TransactionServices
}

func NewTransactionHandler(service interfaces.TransactionServices) *TransactionHandler {
	return &TransactionHandler{TransactionService: service}
}

func (h TransactionHandler) GetAllTransactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactions, err := h.TransactionService.GetAllTransactions()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, ErrTransactionInternalServerError.Error())
			return
		}
		response.JSON(w, http.StatusOK, dto.TransactionResponseList{
			Data:    transactions,
			Code:    strconv.Itoa(http.StatusOK),
			Message: "Find all transactions",
		})
	}
}

func (h TransactionHandler) GetTransactionByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Transaction ID")
			return
		}

		transaction, err := h.TransactionService.GetTransactionByID(userID)
		if err != nil {
			response.Error(w, http.StatusNotFound, ErrTransactionHandlerNotFound.Error())
			return
		}

		response.JSON(w, http.StatusOK, dto.TransactionResponse{
			Data:    transaction,
			Code:    strconv.Itoa(http.StatusOK),
			Message: "Transaction found",
		})
	}
}

func (h TransactionHandler) CreateTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dto.TransactionCreateRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		userId := r.Context().Value("userID").(*uuid.UUID)

		transaction := mappers.ToTransactionModel(request, userId)

		createdTransaction, err := h.TransactionService.CreateTransaction(&transaction)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, ErrTransactionInternalServerError.Error())
			return
		}

		response.JSON(w, http.StatusCreated, dto.TransactionResponse{
			Data:    createdTransaction,
			Code:    strconv.Itoa(http.StatusCreated),
			Message: "Transaction created",
		})
	}
}

func (h TransactionHandler) UpdateTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Transaction ID")
			return
		}

		var request dto.TransactionUpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		transaction := mappers.ToTransactionUpdateModel(request, &userID)

		updatedTransaction, err := h.TransactionService.UpdateTransaction(userID, &transaction)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, ErrTransactionInternalServerError.Error())
			return
		}

		response.JSON(w, http.StatusOK, dto.TransactionResponse{
			Data:    updatedTransaction,
			Code:    strconv.Itoa(http.StatusOK),
			Message: "Transaction updated",
		})
	}
}

func (h TransactionHandler) DeleteTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Transaction ID")
			return
		}

		if err := h.TransactionService.DeleteTransaction(userID); err != nil {
			response.Error(w, http.StatusNotFound, ErrTransactionHandlerNotFound.Error())
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

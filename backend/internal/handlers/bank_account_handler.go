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
	ErrBankAccountHandlerUnprocessableEntity = apperrors.ErrUnprocessableEntity("BankAccount")
	ErrBankAccountHandlerNotFound            = apperrors.ErrNotFound("BankAccount")
	ErrBankAccountHandlerInternalServerError = apperrors.ErrInternalServerError("BankAccount")
)

type BankAccountHandler struct {
	BankAccountService interfaces.BankAccountServices
}

func NewBankAccountHandler(bankAccountService interfaces.BankAccountServices) *BankAccountHandler {
	return &BankAccountHandler{
		BankAccountService: bankAccountService,
	}
}

func (bk BankAccountHandler) GetAllBankAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bankAccount, err := bk.BankAccountService.GetAllBankAccounts()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, ErrBankAccountHandlerInternalServerError.Error())
			return
		}
		response.JSON(w, http.StatusOK, dto.BankAccountResponseList{
			Data:    bankAccount,
			Code:    strconv.Itoa(http.StatusOK),
			Message: "Find all bankAccounts",
		})
	}
}

func (bk BankAccountHandler) GetBankAccountByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Bank account ID")
			return
		}

		bankAccount, err := bk.BankAccountService.GetBankAccountByID(userID)
		if err != nil {
			response.Error(w, http.StatusNotFound, ErrBankAccountHandlerNotFound.Error())
			return
		}

		response.JSON(w, http.StatusOK, dto.BankAccountResponse{
			Data:    bankAccount,
			Code:    strconv.Itoa(http.StatusOK),
			Message: "Bank Account found",
		})
	}
}

func (bk BankAccountHandler) CreateBankAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dto.BankAccountCreateRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		userID := r.Context().Value("userID").(*uuid.UUID)
		bankAccount := mappers.ToBankAccountModel(request, userID)

		createdBankAccount, err := bk.BankAccountService.CreateBankAccount(&bankAccount)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, ErrBankAccountHandlerInternalServerError.Error())
			return
		}

		response.JSON(w, http.StatusCreated, dto.BankAccountResponse{
			Data:    createdBankAccount,
			Code:    strconv.Itoa(http.StatusCreated),
			Message: "BankAccount created",
		})
	}
}

func (bk BankAccountHandler) UpdateBankAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Bank account ID")
			return
		}

		var request dto.BankAccountUpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		bankAccount := mappers.ToBankAccountUpdateModel(request, &userID)
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, ErrBankAccountHandlerUnprocessableEntity.Error())
			return
		}

		updatedBankAccount, err := bk.BankAccountService.UpdateBankAccount(userID, &bankAccount)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, ErrBankAccountHandlerInternalServerError.Error())
			return
		}

		response.JSON(w, http.StatusOK, dto.BankAccountResponse{
			Data:    updatedBankAccount,
			Code:    strconv.Itoa(http.StatusOK),
			Message: "BankAccount updated",
		})
	}
}

func (bk BankAccountHandler) DeleteBankAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Bank account ID")
			return
		}

		if err := bk.BankAccountService.DeleteBankAccount(userID); err != nil {
			response.Error(w, http.StatusNotFound, ErrBankAccountHandlerNotFound.Error())
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

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
	ErrCreditCardHandlerUnprocessableEntity = apperrors.ErrUnprocessableEntity("CreditCard")
	ErrCreditCardHandlerNotFound            = apperrors.ErrNotFound("CreditCard")
	ErrCreditCardHandlerInternalServerError = apperrors.ErrInternalServerError("CreditCard")
)

type CreditCardHandler struct {
	CreditCardService interfaces.CreditCardServices
}

func NewCreditCardHandler(creditCardService interfaces.CreditCardServices) *CreditCardHandler {
	return &CreditCardHandler{CreditCardService: creditCardService}
}

func (cc CreditCardHandler) GetAllCreditCard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("userID").(*uuid.UUID)
		creditCard, err := cc.CreditCardService.GetAllCreditCards(userId)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, ErrCreditCardHandlerInternalServerError.Error())
			return
		}
		response.JSON(w, http.StatusOK, dto.CreditCardResponseList{
			Data:    creditCard,
			Code:    strconv.Itoa(http.StatusOK),
			Message: "Find all creditCards",
		})
	}
}

func (cc CreditCardHandler) GetCreditCardByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		creditCardId, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Credit card ID")
			return
		}
		userId := r.Context().Value("userID").(*uuid.UUID)

		creditCard, err := cc.CreditCardService.GetCreditCardByID(&creditCardId, userId)
		if err != nil {
			response.Error(w, http.StatusNotFound, ErrCreditCardHandlerNotFound.Error())
			return
		}

		response.JSON(w, http.StatusOK, dto.CreditCardResponse{
			Data:    creditCard,
			Code:    strconv.Itoa(http.StatusOK),
			Message: "CreditCard found",
		})
	}
}

func (cc CreditCardHandler) CreateCreditCard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dto.CreditCardCreateRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		userID := r.Context().Value("userID").(*uuid.UUID)
		creditCard := mappers.ToCreditCardModel(request, userID)

		createdCreditCard, err := cc.CreditCardService.CreateCreditCard(&creditCard)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, ErrCreditCardHandlerInternalServerError.Error())
			return
		}

		response.JSON(w, http.StatusCreated, dto.CreditCardResponse{
			Data:    createdCreditCard,
			Code:    strconv.Itoa(http.StatusCreated),
			Message: "CreditCard created",
		})
	}
}

func (cc CreditCardHandler) UpdateCreditCard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Credit Card ID")
			return
		}

		var request dto.CreditCardUpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		userID = r.Context().Value("userID").(uuid.UUID)
		creditCard := mappers.ToCreditCardUpdateModel(request, &userID)
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, ErrCreditCardHandlerUnprocessableEntity.Error())
			return
		}

		updatedCreditCard, err := cc.CreditCardService.UpdateCreditCard(userID, &creditCard)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, ErrCreditCardHandlerInternalServerError.Error())
			return
		}

		response.JSON(w, http.StatusOK, dto.CreditCardResponse{
			Data:    updatedCreditCard,
			Code:    strconv.Itoa(http.StatusOK),
			Message: "CreditCard updated",
		})
	}
}

func (cc CreditCardHandler) DeleteCreditCard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid Credit card ID")
			return
		}

		if err := cc.CreditCardService.DeleteCreditCard(userID); err != nil {
			response.Error(w, http.StatusNotFound, ErrCreditCardHandlerNotFound.Error())
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

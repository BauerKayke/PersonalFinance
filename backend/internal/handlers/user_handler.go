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
	ErrUserHandlerUnprocessableEntity = apperrors.ErrUnprocessableEntity("User")
	ErrUserHandlerNotFound            = apperrors.ErrNotFound("User")
	ErrUserHandlerInternalServerError = apperrors.ErrInternalServerError("User")
)

type UserHandler struct {
	UserService interfaces.UserServices
}

func NewUserHandler(userHandler interfaces.UserServices) *UserHandler {
	return &UserHandler{UserService: userHandler}
}

func (u UserHandler) GetAllUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := u.UserService.GetAllUsers()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, ErrUserHandlerInternalServerError.Error())
			return
		}
		response.JSON(w, http.StatusOK, dto.UserResponseList{
			Data:    users,
			Code:    strconv.Itoa(http.StatusOK),
			Message: "Find all bankAccounts",
		})
	}
}

func (u UserHandler) GetUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid User ID")
			return
		}

		user, err := u.UserService.GetUserByID(userID)
		if err != nil {
			response.Error(w, http.StatusNotFound, ErrUserHandlerNotFound.Error())
			return
		}

		response.JSON(w, http.StatusOK, dto.UserResponse{
			Data:    user,
			Code:    strconv.Itoa(http.StatusOK),
			Message: "Bank Account found",
		})
	}
}

func (u UserHandler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request dto.UserCreateRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		user := mappers.ToUserModel(request)

		createdUser, err := u.UserService.CreateUser(&user)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, ErrUserHandlerInternalServerError.Error())
			return
		}

		response.JSON(w, http.StatusCreated, dto.UserResponse{
			Data:    createdUser,
			Code:    strconv.Itoa(http.StatusCreated),
			Message: "Transaction created",
		})
	}
}

func (u UserHandler) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid User ID")
			return
		}

		var request dto.UserUpdatedRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		user := mappers.ToUserUpdateModel(request)

		updatedUser, err := u.UserService.UpdateUser(userID, &user)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, ErrUserHandlerInternalServerError.Error())
			return
		}

		response.JSON(w, http.StatusOK, dto.UserResponse{
			Data:    updatedUser,
			Code:    strconv.Itoa(http.StatusOK),
			Message: "Transaction updated",
		})
	}
}

func (u UserHandler) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userID, err := uuid.Parse(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "Invalid User ID")
			return
		}

		if err := u.UserService.DeleteUser(userID); err != nil {
			response.Error(w, http.StatusNotFound, ErrUserHandlerNotFound.Error())
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

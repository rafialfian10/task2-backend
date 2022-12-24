package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	authdto "project/dto/auth"
	dto "project/dto/result"
	"project/models"
	"project/pkg/bcrypt"
	jwtToken "project/pkg/jwt"
	"project/repositories"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

// Handle struct
type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

// Handle Authentication
func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

// Function register
func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Hashing password
	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
		Phone:    request.Phone,
		Address:  request.Address,
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseRegister(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseRegister(u models.User) authdto.RegisterResponse {
	return authdto.RegisterResponse{
		Email:    u.Email,
		Password: u.Password,
	}
}

// Function Login
func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.AuthRepository.Login(user.Email)
	// misal email gak ada
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "salah email atau password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["email"] = user.Email
	claims["password"] = user.Password
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, err := jwtToken.GenerateToken(&claims)
	if err != nil {
		log.Println(err)
		fmt.Println("Unauthorize")
		return
	}

	loginResponse := authdto.LoginResponse{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/dto"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/entity"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/httputil"
	"github.com/victor-bologna/pos-curso-go-expert-apis/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

// GenerateToken godoc
//
//	@Summary		Generate JWT
//	@Description	Generate an JWT based on user.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.GenerateTokenDTO	true	"User credentials"
//	@Success		200 {object} 	dto.JWTString
//	@Failure		400	{object}	httputil.Error
//	@Failure		401	{object}	httputil.Error
//	@Failure		404	{object}	httputil.Error
//	@Failure		500	{object}	httputil.Error
//	@Router			/users/generate_token [post]
func (u *UserHandler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("Jwt").(*jwtauth.JWTAuth)
	expiresIn := r.Context().Value("JwtExpiresIn").(int)

	var generateTokenDTO dto.GenerateTokenDTO

	err := json.NewDecoder(r.Body).Decode(&generateTokenDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	user, err := u.UserDB.FindByEmail(generateTokenDTO.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if !user.ValidatePassword(generateTokenDTO.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		err := httputil.Error{Message: "user not found."}
		json.NewEncoder(w).Encode(err)
		return
	}

	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(expiresIn)).Unix(),
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	accessToken := dto.JWTString{AccessToken: tokenString}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
}

// User godoc
//
//	@Summary		Create User
//	@Description	Create a new user.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.UserDTO	true	"User request"
//	@Success		201
//	@Failure		400	{object}	httputil.Error
//	@Failure		500	{object}	httputil.Error
//	@Router			/users [post]
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	user, err := entity.NewUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	err = u.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := httputil.Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

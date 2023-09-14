package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Fadlihardiyanto/go-jwt-mux/config"
	"github.com/Fadlihardiyanto/go-jwt-mux/helper"
	"github.com/Fadlihardiyanto/go-jwt-mux/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// mengambil data dari body
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// mencari user di database
	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]interface{}{
				"status":  "failed",
				"message": "User tidak ditemukan",
			}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		default:
			response := map[string]interface{}{
				"status":  "failed",
				"message": err.Error(),
			}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return

		}
	}

	// membandingkan password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "Password salah",
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// proses generate token jwt
	expTime := time.Now().Add(1 * time.Minute)
	claims := config.JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasikan algorithm yang digunakan untuk signing token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// proses signing token
	signedToken, err := token.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    signedToken,
		HttpOnly: true,
	})

	// kirim response
	response := map[string]interface{}{
		"status":  "success",
		"message": "Berhasil login",
	}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func Register(w http.ResponseWriter, r *http.Request) {

	// mengambil data dari body
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// hass password menggunakan bcrypt
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)

	userInput.Password = string(hashedPassword)

	// simpan data ke database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": err.Error(),
		}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// kirim response

	response := map[string]interface{}{
		"status":  "success",
		"message": "Berhasil register",
	}

	helper.ResponseJSON(w, http.StatusOK, response)

}

func Logout(w http.ResponseWriter, r *http.Request) {

	// hapus cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	// kirim response
	response := map[string]interface{}{
		"status":  "success",
		"message": "Berhasil logout",
	}
	helper.ResponseJSON(w, http.StatusOK, response)
}

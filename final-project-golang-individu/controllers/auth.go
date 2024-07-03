package controllers

import (
	"final-project-golang-individu/config"
	"final-project-golang-individu/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Profile  struct {
		FullName string `json:"full_name" binding:"required"`
		Bio      string `json:"bio"`
	} `json:"profile"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email, password and profile information.
// @Tags auth
// @Accept json
// @Produce json
// @Param register body RegisterInput true "Register Input"
// @Success 200 {object} RegisterResponse
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		Profile: models.Profile{
			FullName: input.Profile.FullName,
			Bio:      input.Profile.Bio,
		},
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, RegisterResponse{Message: "registration success"})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Login godoc
// @Summary Login a user
// @Description Login a user and get JWT token.
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginInput true "Login Input"
// @Success 200 {object} LoginResponse
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{Token: tokenString})
}

type ChangePasswordResponse struct {
	Message string `json:"message"`
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change the password of the authenticated user.
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} ChangePasswordResponse
// @Router /auth/change-password [post]
func ChangePassword(c *gin.Context) {
	// Handle password change
	c.JSON(http.StatusOK, ChangePasswordResponse{Message: "Change Password"})
}

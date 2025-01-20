package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/siddharthTricon/go-task-management-sysytem/services"
	"github.com/siddharthTricon/go-task-management-sysytem/utils"
	"github.com/siddharthTricon/go-task-management-sysytem/models"
	"github.com/siddharthTricon/go-task-management-sysytem/database"
)

type AuthController struct{
	jwtService services.JWTService
}

func NewAuthController(jwtService services.JWTService) *AuthController{
	return &AuthController{jwtService: jwtService}
}

func (ac *AuthController) Register(c *gin.Context){
	var user models.User
	if err := c.BindJSON(&user); err != nil{
		utils.RespondJSON(c, http.StatusBadRequest,"Invalid request data",gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err !=nil{
		utils.RespondJSON(c, http.StatusInternalServerError,"Error ecncrypting password", gin.H{"error":"Error encrypting password"})
		return
	}

	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil{
		utils.RespondJSON(c, http.StatusInternalServerError,"Error saving user", gin.H{"error":"Error saving user"})
		return
	}

	utils.RespondJSON(c, http.StatusCreated,"User successfully created", user)
}

func (ac *AuthController) Login(c *gin.Context){
		var user models.User
		var loginRequest models.User
	
	if err := c.BindJSON(&loginRequest); err != nil{
		utils.RespondJSON(c, http.StatusBadRequest,"Invalid request data", gin.H{"error":err.Error()})
		return
	}

	if err := database.DB.Where("username = ?", loginRequest.Username).First(&user).Error; err != nil{
		utils.RespondJSON(c, http.StatusUnauthorized,"Invalid username and password", gin.H{"error":"Invalid username and password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil{
		utils.RespondJSON(c, http.StatusUnauthorized,"Invalid username and password", gin.H{"error":"Invalid username or password"})
		return
	}

	token, err := ac.jwtService.GenerateToken(user.ID, user.Role)
	if err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError,"Failed to generate token", gin.H{"error":"Failed to generate token"})
		return
	}
	utils.RespondJSON(c, http.StatusOK,"Login Successful", gin.H{"token":token})
}
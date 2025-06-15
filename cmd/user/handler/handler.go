package handler

import (
	"net/http"
	"user/cmd/user/usecase"
	"user/infrastructure/log"
	"user/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	var param models.LoginParameter

	if err := c.ShouldBindJSON(&param); err != nil {
		log.Logger.Info(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"error_messages": "Invalid input parameters.",
		})

		return
	}

	token, err := h.UserUsecase.Login(c.Request.Context(), &param)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_messages": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *UserHandler) Register(c *gin.Context) {
	var param models.RegisterParameter

	if err := c.ShouldBindJSON(&param); err != nil {
		log.Logger.Info(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{
			"error_messages": "Invalid input parameters.",
		})

		return
	}

	// check user availability
	user, err := h.UserUsecase.GetUserByEmail(c.Request.Context(), param.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_messages": err.Error(),
		})

		return
	}

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_messages": "This email is already registered. Try logging in instead.",
		})

		return
	}

	// call register use case
	err = h.UserUsecase.RegisterUser(c.Request.Context(), &models.User{
		Name:     param.Name,
		Email:    param.Email,
		Password: param.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_messages": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User has been registered.",
	})
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	// extract user id from jwt claim
	userIdStr, isExist := c.Get("user_id")

	if !isExist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized.",
		})

		return
	}

	userId, ok := userIdStr.(float64)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid user id.",
		})

		return
	}

	user, err := h.UserUsecase.GetUserById(c.Request.Context(), int64(userId))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found.",
		})

		return
	}

	// valid
	c.JSON(http.StatusOK, gin.H{
		"name":  user.Name,
		"email": user.Email,
	})
}

func (h *UserHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "Ok",
	})
}

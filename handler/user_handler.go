package handler

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type userHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) Register(c *gin.Context) {
	var userRegister model.Register

	if err := c.ShouldBindJSON(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, model.CreateHandlerResponseError(err.Error()))
		return
	}

	if userRegister.Password != userRegister.ConfirmPassword {
		c.JSON(http.StatusBadRequest, model.CreateHandlerResponseError("Password does not match"))
		return
	}

	if err := h.userService.Register(&userRegister); err != nil {
		c.JSON(http.StatusInternalServerError, model.CreateHandlerResponseError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.CreateHandlerResponseSuccess("Register successfully"))

}

func (h *userHandler) Login(c *gin.Context) {
	var userLogin model.Login

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, model.CreateHandlerResponseError(err.Error()))
		return
	}

	token, userID, err := h.userService.Login(&userLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.CreateHandlerResponseError(err.Error()))
		return
	}

	c.SetCookie("session_token", *token, 3600, "/", "localhost", false, false)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login sucess",
		"userId":  userID,
	})
}

func (h *userHandler) Logout(c *gin.Context) {
	c.SetCookie("session_token", "", -1, "/", "", false, false)
	c.JSON(http.StatusOK, model.CreateHandlerResponseSuccess("Logout successfully"))
}

package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @ID register
// @Router /register [POST]
// @Summary Register
// @Description Register
// @Tags Register
// @Accept json
// @Produce json
// @Param register body models.Register true "RegisterRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *handler) Register(c *gin.Context) {
	var registerUser models.Register
	err := c.ShouldBindJSON(&registerUser)
	if err != nil {
		h.handlerResponse(c, "register user bind", http.StatusBadRequest, err.Error())
		return
	}
	if len(registerUser.Login) < 6 {
		log.Println("length of the login must contain more than 6 letters!! ")
	}
	if len(registerUser.Password) < 6 {
		log.Println("length of the password must contain more than 6 letters!! ")
	}
	_, err = h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Login: registerUser.Login})
	if err == nil {
		h.handlerResponse(c, "login check user", http.StatusBadRequest, "user already exists, please login")
		return
	} else {
		if err.Error() == "no rows in result set" {
		} else {
			h.handlerResponse(c, "login check user", http.StatusBadRequest, "err.Error()")
			return
		}
	}
	id, err := h.strg.User().Create(context.Background(), &models.CreateUser{
		FirstName:   registerUser.FirstName,
		LastName:    registerUser.LastName,
		Login:       registerUser.Login,
		Password:    registerUser.Password,
		PhoneNumber: registerUser.PhoneNumber,
	})
	if err != nil {
		h.handlerResponse(c, "create user", http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{
		Id: id,
	})
	if err != nil {
		h.handlerResponse(c, "storage.create.getbyid user", http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Login
// @Description Login
// @Tags Login
// @Accept json
// @Produce json
// @Param login body models.Login true "LoginRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *handler) Login(c *gin.Context) {
	var loginUser models.Login

	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		h.handlerResponse(c, "login user bind", http.StatusBadRequest, err.Error())
		return
	}
	if len(loginUser.Login) < 6 {
		log.Println("length of the login must contain more than 6 letters!! ")
	}
	if len(loginUser.Password) < 6 {
		log.Println("length of the password must contain more than 6 letters!! ")
	}
	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Login: loginUser.Login})
	if err != nil {
		if err.Error() == "no rows in reslut set" {
			h.handlerResponse(c, "login check user", http.StatusBadRequest, " user not found")
			return
		}
		h.handlerResponse(c, "login check user", http.StatusBadRequest, err.Error())
		return
	}

	if loginUser.Password != resp.Password {
		h.handlerResponse(c, "password check", http.StatusBadRequest, "incorect password")
		return
	}
	data := make(map[string]interface{})
	data["user_id"] = resp.Id
	token, err := helper.GenerateJWT(data, time.Hour*100, h.cfg.SecretKey)
	if err != nil {
		h.handlerResponse(c, "token generate", http.StatusBadRequest, "incorect password")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

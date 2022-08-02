package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wikankun/user-service/config"
	"github.com/wikankun/user-service/model"
	"github.com/wikankun/user-service/util"
)

type paramUserGetToken struct {
	Username string
	Password string
}

type respUserGetToken struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire_time"`
}

type paramUserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type respUserRegister struct {
	ID int `json:"id"`
}

type paramUserVerify struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}

type paramUserUpdate struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type paramUserDelete struct {
	ID int `json:"id"`
}

type paramUserInfo struct {
	ID int `json:"id"`
}

func UserGetToken(c echo.Context) error {
	var param paramUserGetToken
	if err := c.Bind(&param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user, err := model.GetUserWithUsername(param.Username)
	if user == (model.User{}) {
		return util.ErrorResponse(c, http.StatusBadRequest, "user not found")
	}
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	if !util.CheckPasswordHash(param.Password, user.Password) {
		return util.ErrorResponse(c, http.StatusForbidden, "username or password don't match")
	}
	if !user.Verified {
		return util.ErrorResponse(c, http.StatusForbidden, "please verify your email")
	}

	token, expireTime, err := util.NewJWTToken(user.ID)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return util.SuccessResponse(c, http.StatusOK, respUserGetToken{
		Token:  token,
		Expire: expireTime.Unix(),
	})
}

func UserRegister(c echo.Context) error {
	var param paramUserRegister
	if err := c.Bind(&param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	user, _ := model.GetUserWithUsername(param.Username)
	if user != (model.User{}) {
		return util.ErrorResponse(c, http.StatusInternalServerError, "username already exist")
	}

	hashedPassword, err := util.HashPassword(param.Password)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	user = model.User{
		Username: param.Username,
		Password: hashedPassword,
		Email:    param.Email,
		Verified: false,
	}
	id, err := model.AddUser(user)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	verificationLength, _ := strconv.Atoi(config.Config.App.VerificationLength)
	verificationCode := util.RandomString(verificationLength)

	err = util.SendEmail(
		param.Email,
		"Email Verification",
		fmt.Sprintf("Verification code: <code>%s</code>", verificationCode),
	)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	err = model.CreateVerification(id, verificationCode)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, http.StatusCreated, respUserRegister{
		ID: id,
	})
}

func UserVerify(c echo.Context) error {
	var param paramUserVerify
	if err := c.Bind(&param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	verification, err := model.VerifyCode(param.Code)
	log.Println(verification)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	user := verification.User
	if user == (model.User{}) {
		return util.ErrorResponse(c, http.StatusBadRequest, "verification code doesn't exist")
	}

	err = model.DeleteVerifyCode(param.Code)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	user.Verified = true

	err = model.UpdateUser(user)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, http.StatusOK, nil)
}

func UserUpdate(c echo.Context) error {
	var param paramUserUpdate
	if err := c.Bind(&param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id := util.MustGetIDFromContext(c)
	user, err := model.GetUserWithID(id)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	info := model.User{}
	if param.Username != "" {
		newUser, err := model.GetUserWithUsername(param.Username)
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		if param.Username != user.Username && newUser != (model.User{}) {
			return util.ErrorResponse(c, http.StatusBadRequest, "username already exists")
		}
		info.Username = param.Username
	}

	if param.Password != "" {
		info.Password, err = util.HashPassword(param.Password)
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
	}

	if param.Email != "" {
		verificationLength, _ := strconv.Atoi(config.Config.App.VerificationLength)
		verificationCode := util.RandomString(verificationLength)
		err = util.SendEmail(
			param.Email,
			"Email Verification",
			fmt.Sprintf("Verification code: <code>%s</code>", verificationCode),
		)
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		err = model.CreateVerification(param.ID, verificationCode)
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		info.Email = param.Email
		info.Verified = false
	}

	err = model.UpdateUser(info)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, http.StatusOK, nil)
}

func UserInfo(c echo.Context) error {
	var param paramUserInfo
	if err := c.Bind(&param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(param); err != nil {
		return util.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	var result []model.User
	id := util.MustGetIDFromContext(c)
	isAdmin, err := model.IsUserAdmin(id)
	if err != nil {
		return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if param.ID != 0 {
		if param.ID != id && !isAdmin {
			return util.ErrorResponse(c, http.StatusForbidden, "you are not admin")
		}

		user, err := model.GetUserWithID(param.ID)
		if user == (model.User{}) {
			return util.ErrorResponse(c, http.StatusBadRequest, "user not found")
		}
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		result = []model.User{user}
	} else {
		if !isAdmin {
			return util.ErrorResponse(c, http.StatusForbidden, "you are not admin")
		}

		users, err := model.GetAllUsers()
		if err != nil {
			return util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		result = users
	}

	response := make([]echo.Map, 0)
	for _, user := range result {
		response = append(response, echo.Map{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"created_at": user.CreatedAt,
			"verified":   user.Verified,
		})
	}

	return util.SuccessResponse(c, http.StatusOK, echo.Map{"result": response})
}

package handler

import (
	"backend-github-trending/model"
	req "backend-github-trending/model/req"
	"backend-github-trending/repository"
	security "backend-github-trending/security"
	"net/http"

	"github.com/labstack/gommon/log"

	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := req.ReqSignUp{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error()) // Changed to log.Println
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error()) // Changed to log.Println
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()

	userId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error()) // Changed to log.Println
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user := model.User{
		UserId:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     role,
		Token:    "",
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		log.Error(err.Error()) // Changed to log.Println
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user.Password = ""
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Processed successfully",
		Data:       user,
	})
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	req := req.ReqSignIn{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error()) // Changed to log.Println
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error()) // Changed to log.Println
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// check pass
	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Đăng nhập thất bại",
			Data:       nil,
		})
	}

	user.Password = ""
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý Thành Công",
		Data:       user,
	})
}

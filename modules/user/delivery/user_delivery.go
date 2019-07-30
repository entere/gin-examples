package delivery

import (
	"log"
	"micro-gin/handler"
	"micro-gin/model"
	"micro-gin/modules/user"
	user_repository "micro-gin/modules/user/repository"
	user_service "micro-gin/modules/user/service"
	user_auth_repository "micro-gin/modules/user_auth/repository"
	user_auth_service "micro-gin/modules/user_auth/service"
	"micro-gin/pkg/auth"
	"micro-gin/pkg/errno"
	"micro-gin/pkg/jwtauth"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
)

// UserDelivery ...
var UserDelivery = newUserDelivery()

type userDelivery struct {
}

func newUserDelivery() *userDelivery {
	return &userDelivery{}
}

// @Summary 用户注册
// @Description 注册
// @Accept  json
// @Produce  json
// @Param   identifier identity_type credential
// @Success 0 {string} string    "ok"
// @Router /signup [post]
func (ud *userDelivery) Signup(c *gin.Context) {
	var (
		userModel       model.User
		userAuthModel   model.UserAuth
		userService     = user_service.NewUserService(user_repository.NewUserRepository())
		userAuthService = user_auth_service.NewUserAuthService(user_auth_repository.NewUserAuthRepository())

		userID = uuid.NewV4().String()
		ID     = uuid.NewV4().String()
	)
	if err := c.Bind(&userAuthModel); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 判断注册用户是否存在
	exists, err := userAuthService.ExistByIdentifierAndIdentityType(userAuthModel.Identifier, userAuthModel.IdentityType)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}
	if exists {
		handler.SendResponse(c, errno.ErrUserExist, nil)
		return
	}

	// 根据用户名类型，判断密码是否需要加密
	if userAuthModel.IdentityType == "mobile" || userAuthModel.IdentityType == "email" || userAuthModel.IdentityType == "username" {
		hashCredential, err := auth.Encrypt(userAuthModel.Credential)
		if err != nil {
			handler.SendResponse(c, errno.ErrEncrypt, nil)
		} else {
			userAuthModel.Credential = hashCredential
		}
	}

	// 添加用户到 user_auths/user表
	userAuthModel.ID = ID
	userAuthModel.UserID = userID
	userModel.ID = userID
	if err := userService.Signup(&userModel, &userAuthModel); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, err)
		return
	}

	res := user.SignupResponse{
		Identifier: userAuthModel.Identifier,
		UserID:     userID,
	}

	// Show the user information.
	handler.SendResponse(c, nil, res)

}

func (ud *userDelivery) Signin(c *gin.Context) {
	var (
		userAuthModel   model.UserAuth
		userAuthService = user_auth_service.NewUserAuthService(user_auth_repository.NewUserAuthRepository())
	)
	if err := c.Bind(&userAuthModel); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 判断注册用户是否存在
	uad, err := userAuthService.GetByIdentifierAndIdentityType(userAuthModel.Identifier, userAuthModel.IdentityType)
	if err != nil {
		handler.SendResponse(c, errno.InternalServerError, nil)
		return
	}

	// 判断密码是否一致
	if err := auth.Compare(uad.Credential, userAuthModel.Credential); err != nil {
		log.Println(err)
		handler.SendResponse(c, errno.ErrPasswordIncorrect, err)
		return
	}

	// Sign the json web token.
	t, err := jwtauth.GenerateToken(c, jwtauth.Claims{UserID: uad.UserID, Username: uad.Identifier}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}

	handler.SendResponse(c, nil, user.SigninResponse{Token: t})

}

func (ud *userDelivery) Show(c *gin.Context) {
	id := c.Param("id")
	userService := user_service.NewUserService(user_repository.NewUserRepository())

	user, err := userService.GetByID(id)
	if err != nil {

		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return

	}

	handler.SendResponse(c, nil, user)
}

package structs

import (
	"errors"
	"money-manager/entity"
	"time"

	"github.com/google/uuid"
)

const (

	MESSAGE_FAILED_GET_DATA_FROM_BODY      = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER           = "failed create user"
	MESSAGE_FAILED_GET_LIST_USER           = "failed get list user"
	MESSAGE_FAILED_GET_USER_TOKEN          = "failed get user token"
	MESSAGE_FAILED_TOKEN_NOT_VALID         = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND         = "token not found"
	MESSAGE_FAILED_GET_USER                = "failed get user"
	MESSAGE_FAILED_LOGIN                   = "failed login"
	MESSAGE_FAILED_WRONG_CREDENTIALS 	   = "wrong credentials"
	MESSAGE_FAILED_UPDATE_USER             = "failed update user"
	MESSAGE_FAILED_DELETE_USER             = "failed delete user"


	MESSAGE_SUCCESS_REGISTER_USER           = "success create user"
	MESSAGE_SUCCESS_GET_LIST_USER           = "success get list user"
	MESSAGE_SUCCESS_GET_USER                = "success get user"
	MESSAGE_SUCCESS_LOGIN                   = "success login"
	MESSAGE_SUCCESS_UPDATE_USER             = "success update user"
	MESSAGE_SUCCESS_DELETE_USER             = "success delete user"
)

var (
	ErrCreateUser             = errors.New("failed to create user")
	ErrGetAllUser             = errors.New("failed to get all user")
	ErrGetUserById            = errors.New("failed to get user by id")
	ErrGetUserByUsername      = errors.New("failed to get user by username")
	ErrUsernameAlreadyExists  = errors.New("username already exist")
	ErrUpdateUser             = errors.New("failed to update user")
	ErrUserNotFound           = errors.New("user not found")
	ErrDeleteUser             = errors.New("failed to delete user")
	ErrCredentials      	  = errors.New("wrong credentials")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrTokenExpired           = errors.New("token expired")
)

type (
	UserCreateRequest struct {
		Username    string	`json:"username" form:"username"`
		Fullname	string	`json:"fullname" form:"fullname"`
		Password	string 	`json:"password" form:"password"`
	}

	UserDataLoginResponse struct {
		UserID      uuid.UUID 	`json:"user_id"`
		Username    string		`json:"username"`
		Fullname	string		`json:"fullname"`
		CreatedAt  	time.Time 	`json:"created_at"`
    	UpdatedAt  	time.Time 	`json:"updated_at"`
	}

	GetAllUserPaginationResponse struct {
		Users []entity.Users
		PaginationResponse
	}

	UserUpdateRequest struct {
		UserID		uuid.UUID	`json:"user_id"`
		Username    string		`json:"username" form:"username"`
		Fullname	string		`json:"fullname" form:"fullname"`
	}

	UserDeletionRequest struct {
		UserID	uuid.UUID	`json:"user_id"`
	}

	UserLoginRequest struct {
		Username    string `json:"username" form:"email" binding:"required"`
		Password 	string `json:"password" form:"password" binding:"required"`
	}

	UserLoginResponse struct {
		Token string `json:"token"`
	}
)
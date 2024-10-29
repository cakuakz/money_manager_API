package controller

import (
	"database/sql"
	"money-manager/repository"
	"money-manager/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(ctx *gin.Context, db *sql.DB) {
    var user structs.UserCreateRequest

    if err := ctx.ShouldBind(&user); err != nil {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
        return
    }

    hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
    }

    user.Password = string(hashedPass)

    if err := repository.RegisterUser(db, user); err != nil {
        ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user", "details": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "you are registered"})
}

func LoginUser(ctx *gin.Context, db *sql.DB) (usr structs.UserLoginRequest, err error) {
    var userPayload structs.UserLoginRequest

    if err := ctx.ShouldBind(&userPayload); err != nil {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
        return usr, err
    }

    userData, err := repository.VerifyUser(ctx, db, userPayload)
    if err != nil {
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"data": userData})
    usr.Username = userData.Username
    usr.Password = userData.Password

    return usr, nil
}
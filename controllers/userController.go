package controllers

import (
	"Rock_Paper_Scossors/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController interface {
	SingUp(ctx *gin.Context)
	GetAllUser(ctx *gin.Context) 
	Login(ctx *gin.Context)
}

type userController struct {
	jwtService service.JWTService
}

func NewLoginController(jwt service.JWTService) UserController {
	return &userController{
		jwtService: jwt,
	}
}

func (controller *userController) SingUp(ctx *gin.Context) {
	// Bind value 
	newUser := service.SingleUser()
	ctx.ShouldBindJSON(newUser)

	// fmt.Println(newUser.Username)


	// Find a user
	user := service.SingleUser()
	// fmt.Println(user)
	_ = mgm.Coll(user).First(bson.M{"username": newUser.Username}, user)

	// if userErr != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"msg": "Some thing went wrong, Please try agin!!!",
	// 	})
	// 	return
	// }

	if user.Username != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "This Username already used",
		})
		return
	}

	// newUser := service.SingleUser()
	err := mgm.Coll(newUser).Create(newUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Can not inset User",
		})
		return
	}

	// generate token
	token := controller.jwtService.GenerateToken(newUser.ID.Hex())

	// hind id
	newUser.ID = primitive.NilObjectID

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "success",
		"data": newUser,
		"token": token,
	})
}

func (controller *userController) GetAllUser(ctx *gin.Context) {
	users := service.ListUser()
	err := mgm.Coll(service.SingleUser()).SimpleFind(users, bson.M{})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Can not view all user",
		})
		return
	} 

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": users,
	})
}

func (controller *userController) Login(ctx *gin.Context) {
	user := service.SingleUser()
	ctx.ShouldBindJSON(user)

	err := mgm.Coll(user).First(bson.M{"username": user.Username}, user)

	if err != nil || user.Username == "" {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "Not found that username",
		})
		return
	}

	// generate token
	token := controller.jwtService.GenerateToken(user.ID.Hex())

	// send token when it's success
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"token": token,
	})

}
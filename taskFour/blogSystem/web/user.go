package web

import (
	"blogSystem/dao"
	. "blogSystem/models"
	"blogSystem/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"log"
	"net/http"
	"strings"
)

type LoginForm struct {
	Username string `json:"username" validate:"required,max=30" `
	Pwd      string `json:"pwd" validate:"required,max=30" `
}

func Login(context *gin.Context) {
	req := LoginForm{}
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, fieldError := range errs {
			//log.Println("fieldError: ", fieldError)
			log.Printf("字段【%s】校验失败，%s", fieldError.Field(), fieldError.Tag())
			context.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("字段【%s】校验失败，%s", fieldError.Field(), fieldError.Tag()),
			})
			return
		}
	}
	user, err := dao.GetUserByUsername(req.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	fmt.Println("user", user)
	user.Pwd = ""
	//设置session
	//session := sessions.Default(context)
	// 生成一个随机的UUID（版本4），作为token
	uuidV4, err := uuid.NewRandom()
	if err != nil {
		fmt.Println("Error generating UUID:", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	// 转换为字符串
	uuidStr := uuidV4.String()
	userKey := strings.ReplaceAll(uuidStr, "-", "")
	fmt.Println("token:", userKey)
	claims, token, err := utils.GenerateToken(&user, userKey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	log.Println("claims:", claims)
	err = utils.RefreshToken(claims)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": token,
	})
	//context.Redirect(http.StatusMovedPermanently, "/index")
}

type RegisterForm struct {
	Username   string  `json:"username" validate:"required,max=30" `
	Pwd        string  `json:"pwd" validate:"required,max=30" `
	PwdConfirm string  `json:"pwdConfirm" validate:"required,eqcsfield=Pwd"` //eqcsfield校验区分大小写，eqfield不区分大小写
	Phone      string  `json:"phone" validate:"required,e164"`               //e164校验手机格式，+86xxxxxxxxxx
	Email      *string `json:"email" validate:"omitempty,email"`             //omitempty允许为空，email校验邮箱格式
}

func ConvertRegisterFormToUser(form RegisterForm) (User, error) {
	var user User
	err := copier.Copy(&user, &form)
	if err != nil {
		return User{}, err
	}
	return user, err
}

func Register(context *gin.Context) {
	form := RegisterForm{}
	err := context.ShouldBindJSON(&form)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	log.Println("form: ", form)
	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, fieldError := range errs {
			//log.Println("fieldError: ", fieldError)
			log.Printf("字段【%s】校验失败，%s", fieldError.Field(), fieldError.Tag())
			context.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("字段【%s】校验失败，%s", fieldError.Field(), fieldError.Tag()),
			})
		}
		return
	}
	hashPwd, err := utils.GenerateHashPassword(form.Pwd)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	form.Pwd = hashPwd
	user, err := ConvertRegisterFormToUser(form)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	fmt.Println("user: ", user)
	err = dao.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": user,
	})
}

func GetUserInfo(context *gin.Context) {
	username := context.Param("username")
	user, err := dao.GetUserByUsername(username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	user.Pwd = ""
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": user,
	})
}

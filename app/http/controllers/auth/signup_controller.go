package auth // 处理用户身份认证相关逻辑
import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"net/http"
)

//注册 控制器

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 请求对象
	type PhoneExistReq struct {
		Phone string `json:"phone"`
	}
	req := PhoneExistReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		// parse error
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// log
		fmt.Println(err.Error())
		return
	}
	// check db data about phone
	c.JSON(http.StatusOK, gin.H{"exist": user.IsPhoneExist(req.Phone)})
}

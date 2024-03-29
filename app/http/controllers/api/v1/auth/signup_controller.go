package auth // 处理用户身份认证相关逻辑
import (
	"github.com/gin-gonic/gin"
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"net/http"
)

//注册 控制器

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	//初始化请求
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}
	//数据查询
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}
	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

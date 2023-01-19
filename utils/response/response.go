package response

import (
	"gin-simple-project/global/my_errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {
	if data == nil || data == "" {
		data = []struct{}{}
	}
	//Context.Header("key2020","value2020")  	//可以根据实际情况在头部添加额外的其他信息
	Context.JSON(httpCode, gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	})
}

// ReturnJsonFromString 将json字符窜以标准json格式返回（例如，从redis读取json格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
	Context.Header("Content-Type", "application/json; charset=utf-8")
	Context.String(httpCode, jsonStr)
}

// 语法糖函数封装

// Success 直接返回成功
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, http.StatusOK, msg, data)
}

// Fail 失败的业务逻辑
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

// Fail 失败的业务逻辑
func ErrorParamsValid(c *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = "参数验证失败"
	}
	ReturnJson(c, http.StatusBadRequest, http.StatusBadRequest, msg, data)
	c.Abort()
	return
}

// ErrorUserBaseInfo 获取请求头中的域账号
func ErrorUserBaseInfo(c *gin.Context) {
	ReturnJson(c, http.StatusBadRequest, http.StatusBadRequest, my_errors.ErrorUserBaseInfo, nil)
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

// ErrorTokenAuthFail 用户权限验证失败
func ErrorTokenAuthFail(c *gin.Context) {
	ReturnJson(c, http.StatusUnauthorized, http.StatusUnauthorized, my_errors.ErrorsNoAuthorization, nil)
	//终止可能已经被加载的其他回调函数的执行
	c.Abort()
}

// ErrorsGetUserFromTaiBai 获取用户信息失败
func ErrorsGetUserFromTaiBai(c *gin.Context) {
	ReturnJson(c, http.StatusInternalServerError, http.StatusInternalServerError, my_errors.ErrorsGetUserFromTaiBai, nil)
	c.Abort()
}

package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MyCode int

type RestResponse struct {
	Code     MyCode      `json:"code"`
	Message  string      `json:"message"`
	Response interface{} `json:"response"`
}

func ResponseOk(c *gin.Context, response interface{}) {
	res := RestResponse{
		Code:     Ok,
		Message:  Ok.Msg(),
		Response: response,
	}
	c.JSON(http.StatusOK, res)
}

func ResponseFailWithCode(c *gin.Context, code MyCode) {
	res := RestResponse{
		Code:     code,
		Message:  code.Msg(),
		Response: nil,
	}
	c.JSON(http.StatusOK, res)
}

func ResponseFailWithCodeMsg(c *gin.Context, code MyCode, msg string) {
	res := RestResponse{
		Code:     code,
		Message:  msg,
		Response: nil,
	}
	c.JSON(http.StatusOK, res)
}

func ResponseFailWithMsg(c *gin.Context, msg string) {
	res := RestResponse{
		Code:     InnerError,
		Message:  msg,
		Response: nil,
	}
	c.JSON(http.StatusOK, res)
}

func ResponseNoLogin(c *gin.Context) {
	res := RestResponse{
		Code:     UNAUTHORIZED,
		Message:  UNAUTHORIZED.Msg(),
		Response: nil,
	}
	c.JSON(http.StatusOK, res)
}
